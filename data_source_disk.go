package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceDiskIdKey = "id"
const DataSourceDiskLabelKey = "label"
const DataSourceDiskPrimaryKey = "primary"
const DataSourceDiskServerIdKey = "server_id"
const DataSourceDiskSizeKey = "size"

// dataSourceDisk() retrieves information about a server's disk.
func dataSourceDisk() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceDiskIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The disk identifier",
				ForceNew:    true,
			},
			DataSourceDiskLabelKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
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
				Computed:    true,
				Description: "The disk size in gigabytes",
			},
		},

		Read: dataSourceDiskRead,
	}
}

// dataSourceDiskRead() reads information about a server's disk.
func dataSourceDiskRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	diskId := d.Get(DataSourceDiskIdKey).(string)
	serverId := d.Get(DataSourceDiskServerIdKey).(string)

	req, reqErr := getClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/disks/%s", serverId, diskId), new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the disk - Reason: The API responded with HTTP %s", res.Status)
	}

	disk := DiskBody{}
	json.NewDecoder(res.Body).Decode(&disk)

	return dataSourceDiskReadResponseBody(d, m, &disk)
}

// dataSourceDiskReadResponseBody() parses information about a server's disk.
func dataSourceDiskReadResponseBody(d *schema.ResourceData, m interface{}, disk *clouddk.DiskBody) error {
	d.SetId(disk.Identifier)

	d.Set(DataSourceDiskLabelKey, disk.Label)
	d.Set(DataSourceDiskPrimaryKey, disk.Primary)
	d.Set(DataSourceDiskSizeKey, disk.Size)

	return nil
}
