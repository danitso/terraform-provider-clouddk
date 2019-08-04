package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceLocationsIdsKey = "ids"
const DataSourceLocationsNamesKey = "names"

// dataSourceLocations() retrieves a list of datacenter locations.
func dataSourceLocations() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceLocationsIdsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourceLocationsNamesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},

		Read: dataSourceLocationsRead,
	}
}

// dataSourceLocationsRead() reads information about datacenter locations.
func dataSourceLocationsRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)
	req, reqErr := clouddk.GetClientRequestObject(&clientSettings, "GET", "locations", new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the locations - Reason: The API responded with HTTP %s", res.Status)
	}

	list := make(clouddk.LocationListBody, 0)
	json.NewDecoder(res.Body).Decode(&list)

	ids := make([]interface{}, len(list))
	names := make([]interface{}, len(list))

	for i, v := range list {
		ids[i] = v.Identifier
		names[i] = v.Name
	}

	d.SetId("locations")

	d.Set(DataSourceLocationsIdsKey, ids)
	d.Set(DataSourceLocationsNamesKey, names)

	return nil
}
