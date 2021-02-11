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
	dataSourceLocationsIdsKey   = "ids"
	dataSourceLocationsNamesKey = "names"
)

// dataSourceLocations retrieves a list of datacenter locations.
func dataSourceLocations() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourceLocationsIdsKey: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			dataSourceLocationsNamesKey: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},

		Read: dataSourceLocationsRead,
	}
}

// dataSourceLocationsRead reads information about datacenter locations.
func dataSourceLocationsRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)
	req, err := clouddk.GetClientRequestObject(&clientSettings, "GET", "locations", new(bytes.Buffer))

	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the locations - Reason: The API responded with HTTP %s", res.Status)
	}

	list := make(clouddk.LocationListBody, 0)
	err = json.NewDecoder(res.Body).Decode(&list)

	if err != nil {
		return err
	}

	ids := make([]interface{}, len(list))
	names := make([]interface{}, len(list))

	for i, v := range list {
		ids[i] = v.Identifier
		names[i] = v.Name
	}

	d.SetId("locations")

	d.Set(dataSourceLocationsIdsKey, ids)
	d.Set(dataSourceLocationsNamesKey, names)

	return nil
}
