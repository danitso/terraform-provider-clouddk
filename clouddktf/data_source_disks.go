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
			dataSourceDisksIDKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
			dataSourceDisksIdsKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceDisksLabelsKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk labels",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceDisksPrimaryKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whether a disk is the primary disk",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			dataSourceDisksSizesKey: {
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
	req, err := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/disks", id), new(bytes.Buffer))

	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the disks - Reason: The API responded with HTTP %s", res.Status)
	}

	disks := clouddk.DiskListBody{}
	err = json.NewDecoder(res.Body).Decode(&disks)

	if err != nil {
		return err
	}

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
