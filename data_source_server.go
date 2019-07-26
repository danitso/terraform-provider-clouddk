package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceServerBootedKey = "booted"
const DataSourceServerCPUsKey = "cpus"
const DataSourceServerDiskIdsKey = "disk_ids"
const DataSourceServerDiskLabelsKey = "disk_labels"
const DataSourceServerDiskPrimaryKey = "disk_primary"
const DataSourceServerDiskSizesKey = "disk_sizes"
const DataSourceServerHostnameKey = "hostname"
const DataSourceServerIdKey = "id"
const DataSourceServerLabelKey = "label"
const DataSourceServerMemoryKey = "memory"

// dataSourceServer() retrieves information about a server.
func dataSourceServer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceServerBootedKey: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the server has been booted",
			},
			DataSourceServerCPUsKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The server's CPU count",
			},
			DataSourceServerDiskIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServerDiskLabelsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk labels",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServerDiskPrimaryKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whether the disk is the primary disk",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			DataSourceServerDiskSizesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk sizes in gigabytes",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			DataSourceServerHostnameKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The server hostname",
			},
			DataSourceServerIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
			DataSourceServerLabelKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The server label",
			},
			DataSourceServerMemoryKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The server's memory allocation in megabytes",
			},
		},

		Read: dataSourceServerRead,
	}
}

// dataSourceServerRead() reads information about a server.
func dataSourceServerRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	id := d.Get(DataSourceServerIdKey).(string)
	req, reqErr := getClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s", id), new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	}

	server := ServerBody{}
	json.NewDecoder(res.Body).Decode(&server)

	diskIds := make([]interface{}, len(server.Disks))
	diskLabels := make([]interface{}, len(server.Disks))
	diskPrimary := make([]interface{}, len(server.Disks))
	diskSizes := make([]interface{}, len(server.Disks))

	for i, v := range server.Disks {
		diskIds[i] = v.Identifier
		diskLabels[i] = v.Label
		diskPrimary[i] = (v.Primary == 1)
		diskSizes[i] = v.Size
	}

	d.SetId(id)

	d.Set(DataSourceServerBootedKey, (server.Booted > 0))
	d.Set(DataSourceServerCPUsKey, server.CPUs)
	d.Set(DataSourceServerDiskIdsKey, diskIds)
	d.Set(DataSourceServerDiskLabelsKey, diskLabels)
	d.Set(DataSourceServerDiskPrimaryKey, diskPrimary)
	d.Set(DataSourceServerDiskSizesKey, diskSizes)
	d.Set(DataSourceServerHostnameKey, server.Hostname)
	d.Set(DataSourceServerLabelKey, server.Label)
	d.Set(DataSourceServerMemoryKey, server.Memory)

	return nil
}
