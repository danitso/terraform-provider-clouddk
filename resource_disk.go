package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform/helper/schema"
)

// resourceDisk() manages a disk.
func resourceDisk() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceDiskLabelKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The disk label",
			},
			DataSourceDiskPrimaryKey: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the disk is the primary disk",
			},
			DataSourceDiskServerIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
			DataSourceDiskSizeKey: &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The disk size in gigabytes",
			},
		},

		Create: resourceDiskCreate,
		Read:   resourceDiskRead,
		Update: resourceDiskUpdate,
		Delete: resourceDiskDelete,
	}
}

// resourceDiskCreate() creates a disk.
func resourceDiskCreate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	serverId := d.Get(DataSourceDiskServerIdKey).(string)

	body := clouddk.DiskCreateBody{
		Label: d.Get(DataSourceDiskLabelKey).(string),
		Size:  d.Get(DataSourceDiskSizeKey).(int),
	}

	reqBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(reqBody).Encode(body)

	if encodeErr != nil {
		return encodeErr
	}

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	lockErr := resourceServerLock(d, m, serverId)

	if lockErr != nil {
		return lockErr
	}

	res, resErr := clouddk.DoClientRequest(&clientSettings, "POST", fmt.Sprintf("cloudservers/%s/disks", serverId), reqBody, []int{200}, 60, 10)

	if resErr != nil {
		resourceServerUnlock(d, m, serverId)

		return resErr
	}

	lockErr = resourceServerUnlock(d, m, serverId)

	if lockErr != nil {
		return lockErr
	}

	disk := clouddk.DiskBody{}
	json.NewDecoder(res.Body).Decode(&disk)

	return dataSourceDiskReadResponseBody(d, m, &disk)
}

// resourceDiskRead reads information about an existing disk.
func resourceDiskRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	diskId := d.Id()
	serverId := d.Get(DataSourceFirewallRuleServerIdKey).(string)

	req, reqErr := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/disks/%s", serverId, diskId), new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	} else if res.StatusCode != 200 {
		if res.StatusCode == 404 {
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Failed to read the disk information - Reason: The API responded with HTTP %s", res.Status)
	}

	disk := clouddk.DiskBody{}
	json.NewDecoder(res.Body).Decode(&disk)

	return dataSourceDiskReadResponseBody(d, m, &disk)
}

// resourceDiskUpdate updates an existing disk.
func resourceDiskUpdate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	serverId := d.Get(DataSourceDiskServerIdKey).(string)
	diskId := d.Id()

	body := clouddk.DiskCreateBody{
		Label: d.Get(DataSourceDiskLabelKey).(string),
		Size:  d.Get(DataSourceDiskSizeKey).(int),
	}

	reqBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(reqBody).Encode(body)

	if encodeErr != nil {
		return encodeErr
	}

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	lockErr := resourceServerLock(d, m, serverId)

	if lockErr != nil {
		return lockErr
	}

	res, resErr := clouddk.DoClientRequest(&clientSettings, "PUT", fmt.Sprintf("cloudservers/%s/disks/%s", serverId, diskId), new(bytes.Buffer), []int{200}, 60, 10)

	if resErr != nil {
		resourceServerUnlock(d, m, serverId)

		return resErr
	}

	lockErr = resourceServerUnlock(d, m, serverId)

	if lockErr != nil {
		return lockErr
	}

	disk := clouddk.DiskBody{}
	json.NewDecoder(res.Body).Decode(&disk)

	return dataSourceDiskReadResponseBody(d, m, &disk)
}

// resourceDiskDelete deletes an existing disk.
func resourceDiskDelete(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	serverId := d.Get(DataSourceDiskServerIdKey).(string)
	diskId := d.Id()

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	lockErr := resourceServerLock(d, m, serverId)

	if lockErr != nil {
		return lockErr
	}

	_, err := clouddk.DoClientRequest(&clientSettings, "DELETE", fmt.Sprintf("cloudservers/%s/disks/%s", serverId, diskId), new(bytes.Buffer), []int{200, 404}, 60, 10)

	if err != nil {
		resourceServerUnlock(d, m, serverId)

		return err
	}

	lockErr = resourceServerUnlock(d, m, serverId)

	if lockErr != nil {
		return lockErr
	}

	d.SetId("")

	return nil
}
