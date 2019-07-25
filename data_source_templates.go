package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceTemplatesIdentifier = "identifier"
const DataSourceTemplatesName = "name"
const DataSourceTemplatesResult = "result"

// dataSourceTemplates() retrieves a list of OS templates.
func dataSourceTemplates() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceTemplatesResult: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						DataSourceTemplatesIdentifier: &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The template identifier",
						},
						DataSourceTemplatesName: &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The template name",
						},
					},
				},
			},
		},

		Read: dataSourceTemplatesRead,
	}
}

// dataSourceTemplatesRead() reads information about OS templates.
func dataSourceTemplatesRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)
	req, reqErr := getClientRequestObject(&clientSettings, "GET", "templates?per-page=1000", new(bytes.Buffer))

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

		locationMap[DataSourceTemplatesIdentifier] = v.Identifier
		locationMap[DataSourceTemplatesName] = v.Name

		result[i] = locationMap
	}

	d.SetId("templates")
	d.Set(DataSourceTemplatesResult, result)

	return nil
}
