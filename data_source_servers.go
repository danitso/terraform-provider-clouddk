/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

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

const (
	dataSourceServersFilterKey         = "filter"
	dataSourceServersFilterHostnameKey = "hostname"
	dataSourceServersHostnamesKey      = "hostnames"
	dataSourceServersIdsKey            = "ids"
	dataSourceServersLabelsKey         = "labels"
	dataSourceServersLocationIdsKey    = "location_ids"
	dataSourceServersLocationNamesKey  = "location_names"
	dataSourceServersPackageIdsKey     = "package_ids"
	dataSourceServersPackageNamesKey   = "package_names"
	dataSourceServersTemplateIdsKey    = "template_ids"
	dataSourceServersTemplateNamesKey  = "template_names"
)

// dataSourceServers retrieves a list of servers.
func dataSourceServers() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourceServersFilterKey: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dataSourceServersFilterHostnameKey: &schema.Schema{
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
			dataSourceServersHostnamesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServersIdsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServersLabelsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServersLocationIdsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServersLocationNamesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServersPackageIdsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServersPackageNamesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServersTemplateIdsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServersTemplateNamesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},

		Read: dataSourceServersRead,
	}
}

// dataSourceServersRead reads information about servers.
func dataSourceServersRead(d *schema.ResourceData, m interface{}) error {
	filter := d.Get(dataSourceServersFilterKey).([]interface{})
	filterHostname := ""

	if len(filter) > 0 {
		filterData := filter[0].(map[string]interface{})
		filterHostname = filterData[dataSourceServersFilterHostnameKey].(string)
	}

	// Prepare the relative path based on the filters.
	path := "cloudservers?per-page=1000"

	if len(filterHostname) > 0 {
		path = fmt.Sprintf("%s&hostname=%s", path, url.QueryEscape(filterHostname))
	}

	// Retrieve the list of templates by invoking the API action.
	clientSettings := m.(clouddk.ClientSettings)
	req, err := clouddk.GetClientRequestObject(&clientSettings, "GET", path, new(bytes.Buffer))

	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the servers - Reason: The API responded with HTTP %s", res.Status)
	}

	list := make(clouddk.ServerListBody, 0)
	err = json.NewDecoder(res.Body).Decode(&list)

	if err != nil {
		return err
	}

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

	d.Set(dataSourceServersHostnamesKey, hostnames)
	d.Set(dataSourceServersIdsKey, ids)
	d.Set(dataSourceServersLabelsKey, labels)
	d.Set(dataSourceServersLocationIdsKey, locationIds)
	d.Set(dataSourceServersLocationNamesKey, locationNames)
	d.Set(dataSourceServersPackageIdsKey, packageIds)
	d.Set(dataSourceServersPackageNamesKey, packageNames)
	d.Set(dataSourceServersTemplateIdsKey, templateIds)
	d.Set(dataSourceServersTemplateNamesKey, templateNames)

	return nil
}
