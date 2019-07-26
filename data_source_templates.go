package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceTemplatesFilterKey = "filter"
const DataSourceTemplatesFilterNameKey = "name"
const DataSourceTemplatesIdsKey = "ids"
const DataSourceTemplatesNamesKey = "names"

// dataSourceTemplates() retrieves a list of OS templates.
func dataSourceTemplates() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceTemplatesFilterKey: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						DataSourceTemplatesFilterNameKey: &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The name filter",
							ForceNew:    true,
						},
					},
				},
				MaxItems: 1,
			},
			DataSourceTemplatesIdsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourceTemplatesNamesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},

		Read: dataSourceTemplatesRead,
	}
}

// dataSourceTemplatesRead() reads information about OS templates.
func dataSourceTemplatesRead(d *schema.ResourceData, m interface{}) error {
	filter := d.Get(DataSourceTemplatesFilterKey).([]interface{})
	filterName := ""

	if len(filter) > 0 {
		filterData := filter[0].(map[string]interface{})
		filterName = filterData[DataSourceTemplatesFilterNameKey].(string)
	}

	// Prepare the relative path based on the filters.
	path := "templates?per-page=1000"

	if len(filterName) > 0 {
		path = fmt.Sprintf("%s&name=%s", path, url.QueryEscape(filterName))
	}

	// Retrieve the list of templates by invoking the API action.
	clientSettings := m.(ClientSettings)
	req, reqErr := getClientRequestObject(&clientSettings, "GET", path, new(bytes.Buffer))

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

	d.SetId("templates")

	d.Set(DataSourceLocationsIdsKey, ids)
	d.Set(DataSourceLocationsNamesKey, names)

	return nil
}
