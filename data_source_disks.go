package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceDisksIdKey = "id"
const DataSourceDisksIdsKey = "ids"
const DataSourceDisksLabelsKey = "labels"
const DataSourceDisksPrimaryKey = "primary"
const DataSourceDisksSizesKey = "sizes"

// dataSourceDisks() retrieves information about a server's disks.
func dataSourceDisks() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceDisksIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
			DataSourceDisksIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceDisksLabelsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk labels",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceDisksPrimaryKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whether a disk is the primary disk",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			DataSourceDisksSizesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk sizes in gigabytes",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},

		Read: dataSourceDisksRead,
	}
}

// dataSourceDisksRead() reads information about a server's disks.
func dataSourceDisksRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	id := d.Get(DataSourceDisksIdKey).(string)
	req, reqErr := getClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/disks", id), new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the disks - Reason: The API responded with HTTP %s", res.Status)
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
		diskPrimary[i] = v.Primary
		diskSizes[i] = v.Size
	}

	d.SetId(id)

	d.Set(DataSourceDisksIdsKey, diskIds)
	d.Set(DataSourceDisksLabelsKey, diskLabels)
	d.Set(DataSourceDisksPrimaryKey, diskPrimary)
	d.Set(DataSourceDisksSizesKey, diskSizes)

	return nil
}
