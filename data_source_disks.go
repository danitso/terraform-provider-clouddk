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
	dataSourceDisksIDKey      = "id"
	dataSourceDisksIdsKey     = "ids"
	dataSourceDisksLabelsKey  = "labels"
	dataSourceDisksPrimaryKey = "primary"
	dataSourceDisksSizesKey   = "sizes"
)

// dataSourceDisks retrieves information about a server's disks.
func dataSourceDisks() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourceDisksIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
			dataSourceDisksIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceDisksLabelsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk labels",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceDisksPrimaryKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whether a disk is the primary disk",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			dataSourceDisksSizesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk sizes in gigabytes",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},

		Read: dataSourceDisksRead,
	}
}

// dataSourceDisksRead reads information about a server's disks.
func dataSourceDisksRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	id := d.Get(dataSourceDisksIDKey).(string)
	req, reqErr := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/disks", id), new(bytes.Buffer))

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

	disks := clouddk.DiskListBody{}
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

	d.Set(dataSourceDisksIdsKey, diskIds)
	d.Set(dataSourceDisksLabelsKey, diskLabels)
	d.Set(dataSourceDisksPrimaryKey, diskPrimary)
	d.Set(dataSourceDisksSizesKey, diskSizes)

	return nil
}
