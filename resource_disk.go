package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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
	clientSettings := m.(ClientSettings)

	serverId := d.Get(DataSourceDiskServerIdKey).(string)

	body := DiskCreateBody{
		Label: d.Get(DataSourceDiskLabelKey).(string),
		Size:  d.Get(DataSourceDiskSizeKey).(int),
	}

	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(body)

	req, reqErr := getClientRequestObject(&clientSettings, "POST", fmt.Sprintf("cloudservers/%s/disks", serverId), reqBody)

	if reqErr != nil {
		return reqErr
	}

	req.Header.Set("Content-Type", "application/json")

	res, resErr := doClientRequest(req, []int{200}, 60, 10)

	if resErr != nil {
		return resErr
	}

	disk := DiskBody{}
	json.NewDecoder(res.Body).Decode(&disk)

	return dataSourceDiskReadResponseBody(d, m, &disk)
}

// resourceDiskRead reads information about an existing disk.
func resourceDiskRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	diskId := d.Id()
	serverId := d.Get(DataSourceFirewallRuleServerIdKey).(string)

	req, reqErr := getClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/disks/%s", serverId, diskId), new(bytes.Buffer))

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

	disk := DiskBody{}
	json.NewDecoder(res.Body).Decode(&disk)

	return dataSourceDiskReadResponseBody(d, m, &disk)
}

// resourceDiskUpdate updates an existing disk.
func resourceDiskUpdate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	diskId := d.Id()
	serverId := d.Get(DataSourceFirewallRuleServerIdKey).(string)

	body := DiskCreateBody{
		Label: d.Get(DataSourceDiskLabelKey).(string),
		Size:  d.Get(DataSourceDiskSizeKey).(int),
	}

	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(body)

	req, reqErr := getClientRequestObject(&clientSettings, "PUT", fmt.Sprintf("cloudservers/%s/disks/%s", serverId, diskId), new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	req.Header.Set("Content-Type", "application/json")

	res, resErr := doClientRequest(req, []int{200}, 60, 10)

	if resErr != nil {
		return resErr
	}

	disk := DiskBody{}
	json.NewDecoder(res.Body).Decode(&disk)

	return dataSourceDiskReadResponseBody(d, m, &disk)
}

// resourceDiskDelete deletes an existing disk.
func resourceDiskDelete(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	diskId := d.Id()
	serverId := d.Get(DataSourceFirewallRuleServerIdKey).(string)

	req, reqErr := getClientRequestObject(&clientSettings, "DELETE", fmt.Sprintf("cloudservers/%s/disks/%s", serverId, diskId), new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	_, err := doClientRequest(req, []int{200, 404}, 60, 10)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
