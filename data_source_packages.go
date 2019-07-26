package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourcePackagesIdsKey = "ids"
const DataSourcePackagesNamesKey = "names"

// dataSourcePackages() retrieves a list of server packages.
func dataSourcePackages() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourcePackagesIdsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourcePackagesNamesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},

		Read: dataSourcePackagesRead,
	}
}

// dataSourcePackagesRead() reads information about server packages.
func dataSourcePackagesRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	req, reqErr := getClientRequestObject(&clientSettings, "GET", "cloudservers/get-packages", new(bytes.Buffer))

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

	ids := make([]interface{}, len(list))
	names := make([]interface{}, len(list))

	for i, v := range list {
		ids[i] = v.Identifier
		names[i] = v.Name
	}

	d.SetId("locations")

	d.Set(DataSourcePackagesIdsKey, ids)
	d.Set(DataSourcePackagesNamesKey, names)

	return nil
}
