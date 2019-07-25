package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceLocationsIdentifier = "identifier"
const DataSourceLocationsName = "name"
const DataSourceLocationsResult = "result"

// dataSourceLocations() retrieves a list of datacenter locations.
func dataSourceLocations() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceLocationsResult: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						DataSourceLocationsIdentifier: &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The location identifier",
						},
						DataSourceLocationsName: &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The location name",
						},
					},
				},
			},
		},

		Read: dataSourceLocationsRead,
	}
}

// dataSourceLocationsRead() reads information about datacenter locations.
func dataSourceLocationsRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	req, reqErr := getClientRequestObject(&clientSettings, "GET", "locations", new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	}

	list := make(LocationListBody, 0)
	json.NewDecoder(res.Body).Decode(&list)

	result := make([]interface{}, len(list))

	for i, v := range list {
		locationMap := make(map[string]interface{})

		locationMap[DataSourceLocationsIdentifier] = v.Identifier
		locationMap[DataSourceLocationsName] = v.Name

		result[i] = locationMap
	}

	d.SetId("locations")
	d.Set(DataSourceLocationsResult, result)

	return nil
}
