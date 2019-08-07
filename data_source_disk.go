package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	dataSourceDiskIDKey       = "id"
	dataSourceDiskLabelKey    = "label"
	dataSourceDiskPrimaryKey  = "primary"
	dataSourceDiskServerIDKey = "server_id"
	dataSourceDiskSizeKey     = "size"
)

// dataSourceDisk retrieves information about a server's disk.
func dataSourceDisk() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourceDiskIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The disk identifier",
				ForceNew:    true,
			},
			dataSourceDiskLabelKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The disk label",
			},
			dataSourceDiskPrimaryKey: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the disk is the primary disk",
			},
			dataSourceDiskServerIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
			dataSourceDiskSizeKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The disk size in gigabytes",
			},
		},

		Read: dataSourceDiskRead,
	}
}

// dataSourceDiskRead reads information about a server's disk.
func dataSourceDiskRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	diskID := d.Get(dataSourceDiskIDKey).(string)
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
		return fmt.Errorf("Failed to read the information about the disk - Reason: The API responded with HTTP %s", res.Status)
	}

	disk := clouddk.DiskBody{}
	err = json.NewDecoder(res.Body).Decode(&disk)

	if err != nil {
		return err
	}

	return dataSourceDiskReadResponseBody(d, m, &disk)
}

// dataSourceDiskReadResponseBody parses information about a server's disk.
func dataSourceDiskReadResponseBody(d *schema.ResourceData, m interface{}, disk *clouddk.DiskBody) error {
	d.SetId(disk.Identifier)

	d.Set(dataSourceDiskLabelKey, disk.Label)
	d.Set(dataSourceDiskPrimaryKey, disk.Primary)
	d.Set(dataSourceDiskSizeKey, disk.Size)

	return nil
}
