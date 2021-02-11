/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package clouddktf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceDisk manages a disk.
func resourceDisk() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourceDiskLabelKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The disk label",
			},
			dataSourceDiskPrimaryKey: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the disk is the primary disk",
			},
			dataSourceDiskServerIDKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
			dataSourceDiskSizeKey: {
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

// resourceDiskCreate creates a disk.
func resourceDiskCreate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	serverID := d.Get(dataSourceDiskServerIDKey).(string)

	body := clouddk.DiskCreateBody{
		Label: d.Get(dataSourceDiskLabelKey).(string),
		Size:  clouddk.CustomInt(d.Get(dataSourceDiskSizeKey).(int)),
	}

	reqBody := new(bytes.Buffer)
	err := json.NewEncoder(reqBody).Encode(body)

	if err != nil {
		return err
	}

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	err = resourceServerLock(d, m, serverID)

	if err != nil {
		return err
	}

	res, err := clouddk.DoClientRequest(&clientSettings, "POST", fmt.Sprintf("cloudservers/%s/disks", serverID), reqBody, []int{200}, 60, 10)

	if err != nil {
		resourceServerUnlock(d, m, serverID)

		return err
	}

	err = resourceServerUnlock(d, m, serverID)

	if err != nil {
		return err
	}

	disk := clouddk.DiskBody{}
	err = json.NewDecoder(res.Body).Decode(&disk)

	if err != nil {
		return err
	}

	return dataSourceDiskReadResponseBody(d, m, &disk)
}

// resourceDiskRead reads information about an existing disk.
func resourceDiskRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	diskID := d.Id()
	serverID := d.Get(dataSourceDiskServerIDKey).(string)

	req, err := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/disks/%s", serverID, diskID), new(bytes.Buffer))

	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		if res.StatusCode == 404 {
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Failed to read the disk information - Reason: The API responded with HTTP %s", res.Status)
	}

	disk := clouddk.DiskBody{}
	err = json.NewDecoder(res.Body).Decode(&disk)

	if err != nil {
		return err
	}

	return dataSourceDiskReadResponseBody(d, m, &disk)
}

// resourceDiskUpdate updates an existing disk.
func resourceDiskUpdate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	diskID := d.Id()
	serverID := d.Get(dataSourceDiskServerIDKey).(string)

	body := clouddk.DiskCreateBody{
		Label: d.Get(dataSourceDiskLabelKey).(string),
		Size:  clouddk.CustomInt(d.Get(dataSourceDiskSizeKey).(int)),
	}

	reqBody := new(bytes.Buffer)
	err := json.NewEncoder(reqBody).Encode(body)

	if err != nil {
		return err
	}

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	err = resourceServerLock(d, m, serverID)

	if err != nil {
		return err
	}

	res, resErr := clouddk.DoClientRequest(&clientSettings, "PUT", fmt.Sprintf("cloudservers/%s/disks/%s", serverID, diskID), new(bytes.Buffer), []int{200}, 60, 10)

	if resErr != nil {
		resourceServerUnlock(d, m, serverID)

		return resErr
	}

	err = resourceServerUnlock(d, m, serverID)

	if err != nil {
		return err
	}

	disk := clouddk.DiskBody{}
	err = json.NewDecoder(res.Body).Decode(&disk)

	if err != nil {
		return err
	}

	return dataSourceDiskReadResponseBody(d, m, &disk)
}

// resourceDiskDelete deletes an existing disk.
func resourceDiskDelete(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	diskID := d.Id()
	serverID := d.Get(dataSourceDiskServerIDKey).(string)

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	err := resourceServerLock(d, m, serverID)

	if err != nil {
		return err
	}

	_, err = clouddk.DoClientRequest(&clientSettings, "DELETE", fmt.Sprintf("cloudservers/%s/disks/%s", serverID, diskID), new(bytes.Buffer), []int{200, 404}, 60, 10)

	if err != nil {
		resourceServerUnlock(d, m, serverID)

		return err
	}

	err = resourceServerUnlock(d, m, serverID)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
