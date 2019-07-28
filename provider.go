package main

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
)

const ProviderConfigurationEndpoint = "endpoint"
const ProviderConfigurationKey = "key"

// Provider() returns the object for this provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc: providerConfigure,
		DataSourcesMap: map[string]*schema.Resource{
			"clouddk_disk":               dataSourceDisk(),
			"clouddk_disks":              dataSourceDisks(),
			"clouddk_firewall_rule":      dataSourceFirewallRule(),
			"clouddk_firewall_rules":     dataSourceFirewallRules(),
			"clouddk_ip_addresses":       dataSourceIPAddresses(),
			"clouddk_locations":          dataSourceLocations(),
			"clouddk_network_interface":  dataSourceNetworkInterface(),
			"clouddk_network_interfaces": dataSourceNetworkInterfaces(),
			"clouddk_packages":           dataSourcePackages(),
			"clouddk_server":             dataSourceServer(),
			"clouddk_servers":            dataSourceServers(),
			"clouddk_templates":          dataSourceTemplates(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"clouddk_firewall_rule": resourceFirewallRule(),
			"clouddk_server":        resourceServer(),
		},
		Schema: map[string]*schema.Schema{
			ProviderConfigurationEndpoint: &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "https://api.cloud.dk/v1",
				Description: "The API endpoint",
			},
			ProviderConfigurationKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The API key",
			},
		},
	}
}

// providerConfigure() configures the provider before processing any IronMQ resources.
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get(ProviderConfigurationEndpoint).(string)

	if len(endpoint) < 1 {
		return nil, errors.New("The API endpoint cannot be an empty string")
	}

	key := d.Get(ProviderConfigurationKey).(string)

	if len(key) < 1 {
		return nil, errors.New("The API key cannot be an empty string")
	}

	clientSettings := ClientSettings{
		Endpoint: endpoint,
		Key:      key,
	}

	return clientSettings, nil
}
