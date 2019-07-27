package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceServerDisksIdKey = "id"
const DataSourceServerDisksIdsKey = "ids"
const DataSourceServerDisksLabelsKey = "labels"
const DataSourceServerDisksPrimaryKey = "primary"
const DataSourceServerDisksSizesKey = "sizes"

// dataSourceServerDisks() retrieves information about a server's disks.
func dataSourceServerDisks() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceServerDisksIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
			DataSourceServerDisksIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServerDisksLabelsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk labels",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServerDisksPrimaryKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whether the disk is the primary disk",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			DataSourceServerDisksSizesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk sizes in gigabytes",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},

		Read: dataSourceServerDisksRead,
	}
}

// dataSourceServerDisksRead() reads information about a server's disks.
func dataSourceServerDisksRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	id := d.Get(DataSourceServerDisksIdKey).(string)
	req, reqErr := getClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/disks", id), new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	}

	disks := DiskListBody{}
	json.NewDecoder(res.Body).Decode(&disks)

	diskIds := make([]interface{}, len(disks))
	diskLabels := make([]interface{}, len(disks))
	diskPrimary := make([]interface{}, len(disks))
	diskSizes := make([]interface{}, len(disks))

	for i, v := range disks {
		diskIds[i] = v.Identifier
		diskLabels[i] = v.Label
		diskPrimary[i] = (v.Primary == 1)
		diskSizes[i] = v.Size
	}

	d.SetId(id)

	d.Set(DataSourceServerDisksIdsKey, diskIds)
	d.Set(DataSourceServerDisksLabelsKey, diskLabels)
	d.Set(DataSourceServerDisksPrimaryKey, diskPrimary)
	d.Set(DataSourceServerDisksSizesKey, diskSizes)

	return nil
}
