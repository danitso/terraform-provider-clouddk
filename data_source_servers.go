package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceServersFilterKey = "filter"
const DataSourceServersFilterHostnameKey = "hostname"
const DataSourceServersHostnamesKey = "hostnames"
const DataSourceServersIdsKey = "ids"
const DataSourceServersLabelsKey = "labels"
const DataSourceServersLocationIdsKey = "location_ids"
const DataSourceServersLocationNamesKey = "location_names"
const DataSourceServersPackageIdsKey = "package_ids"
const DataSourceServersPackageNamesKey = "package_names"
const DataSourceServersTemplateIdsKey = "template_ids"
const DataSourceServersTemplateNamesKey = "template_names"

// dataSourceServers() retrieves a list of servers.
func dataSourceServers() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceServersFilterKey: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						DataSourceServersFilterHostnameKey: &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The hostname filter",
							ForceNew:    true,
						},
					},
				},
				MaxItems: 1,
			},
			DataSourceServersHostnamesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServersIdsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServersLabelsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServersLocationIdsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServersLocationNamesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServersPackageIdsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServersPackageNamesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServersTemplateIdsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServersTemplateNamesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},

		Read: dataSourceServersRead,
	}
}

// dataSourceServersRead() reads information about servers.
func dataSourceServersRead(d *schema.ResourceData, m interface{}) error {
	filter := d.Get(DataSourceServersFilterKey).([]interface{})
	filterHostname := ""

	if len(filter) > 0 {
		filterData := filter[0].(map[string]interface{})
		filterHostname = filterData[DataSourceServersFilterHostnameKey].(string)
	}

	// Prepare the relative path based on the filters.
	path := "cloudservers?per-page=1000"

	if len(filterHostname) > 0 {
		path = fmt.Sprintf("%s&hostname=%s", path, url.QueryEscape(filterHostname))
	}

	// Retrieve the list of templates by invoking the API action.
	clientSettings := m.(clouddk.ClientSettings)
	req, reqErr := clouddk.GetClientRequestObject(&clientSettings, "GET", path, new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the servers - Reason: The API responded with HTTP %s", res.Status)
	}

	list := make(clouddk.ServerListBody, 0)
	json.NewDecoder(res.Body).Decode(&list)

	hostnames := make([]interface{}, len(list))
	ids := make([]interface{}, len(list))
	labels := make([]interface{}, len(list))
	locationIds := make([]interface{}, len(list))
	locationNames := make([]interface{}, len(list))
	packageIds := make([]interface{}, len(list))
	packageNames := make([]interface{}, len(list))
	templateIds := make([]interface{}, len(list))
	templateNames := make([]interface{}, len(list))

	for i, v := range list {
		hostnames[i] = v.Hostname
		ids[i] = v.Identifier
		labels[i] = v.Label
		locationIds[i] = v.Location.Identifier
		locationNames[i] = v.Location.Name
		packageIds[i] = v.Package.Identifier
		packageNames[i] = v.Package.Name
		templateIds[i] = v.Template.Identifier
		templateNames[i] = v.Template.Name
	}

	d.SetId("servers")

	d.Set(DataSourceServersHostnamesKey, hostnames)
	d.Set(DataSourceServersIdsKey, ids)
	d.Set(DataSourceServersLabelsKey, labels)
	d.Set(DataSourceServersLocationIdsKey, locationIds)
	d.Set(DataSourceServersLocationNamesKey, locationNames)
	d.Set(DataSourceServersPackageIdsKey, packageIds)
	d.Set(DataSourceServersPackageNamesKey, packageNames)
	d.Set(DataSourceServersTemplateIdsKey, templateIds)
	d.Set(DataSourceServersTemplateNamesKey, templateNames)

	return nil
}
