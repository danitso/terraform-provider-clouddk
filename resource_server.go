package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

const ResourceServerHostname = "hostname"
const ResourceServerLabel = "label"
const ResourceServerLocationId = "location_id"
const ResourceServerPackageId = "package_id"
const ResourceServerRootPassword = "root_password"
const ResourceServerTemplateId = "template_id"

// resourceServer() manages a server.
func resourceServer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			ResourceServerHostname: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The hostname",
			},
			ResourceServerLabel: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The label",
			},
			ResourceServerLocationId: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The location identifier",
				ForceNew:    true,
			},
			ResourceServerPackageId: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The package identifier",
				ForceNew:    true,
			},
			ResourceServerRootPassword: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The root password",
				ForceNew:    true,
				Sensitive:   true,
			},
			ResourceServerTemplateId: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The template identifier",
				ForceNew:    true,
			},
			DataSourceServerBootedKey: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the server has been booted",
			},
			DataSourceServerCPUsKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The server's CPU count",
			},
			DataSourceServerDiskIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServerDiskLabelsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk labels",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServerDiskPrimaryKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whether a disk is the primary disk",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			DataSourceServerDiskSizesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk sizes in gigabytes",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			DataSourceServerLocationNameKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The location name",
			},
			DataSourceServerMemoryKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The server's memory allocation in megabytes",
			},
			DataSourceServerNetworkInterfaceAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The IP addresses assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceDefaultFirewallRulesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The default firewall rules for the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServerNetworkInterfaceFirewallRulesAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The CIDR blocks for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceFirewallRulesCommandsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The commands for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceFirewallRulesIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The identifiers for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceFirewallRulesPortsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The ports of the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceFirewallRulesProtocolsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The protocols for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceGatewaysKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The gateways assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's network interface identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServerNetworkInterfaceLabelsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's network interface labels",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServerNetworkInterfaceNetmasksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The netmasks assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceNetworksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The networks assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfacePrimaryKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whether a network interface is the primary interface",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			DataSourceServerNetworkInterfaceRateLimitsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The rate limits for the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			DataSourceServerPackageNameKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The package name",
			},
			DataSourceServerTemplateNameKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The template name",
			},
		},

		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,
	}
}

// resourceServerCreate() creates a server.
func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	body := ServerCreateBody{
		Hostname:            d.Get(ResourceServerHostname).(string),
		Label:               d.Get(ResourceServerLabel).(string),
		InitialRootPassword: d.Get(ResourceServerRootPassword).(string),
		Package:             d.Get(ResourceServerPackageId).(string),
		Template:            d.Get(ResourceServerTemplateId).(string),
		Location:            d.Get(ResourceServerLocationId).(string),
	}

	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(body)

	res, resErr := doClientRequest(&clientSettings, "POST", "cloudservers", reqBody, []int{200}, 60, 10)

	if resErr != nil {
		return resErr
	}

	server := ServerBody{}
	json.NewDecoder(res.Body).Decode(&server)

	parseErr := dataSourceServerReadResponseBody(d, m, &server)

	if parseErr != nil {
		return parseErr
	}

	if d.Get(DataSourceServerBootedKey).(bool) {
		return nil
	}

	// Keep attempting to boot the server to bypass an API issue which causes the booted flag to remain 0 (false).
	res, resErr = doClientRequest(&clientSettings, "POST", fmt.Sprintf("cloudservers/%s/start", server.Identifier), new(bytes.Buffer), []int{200}, 20, 30)

	if resErr != nil {
		return resErr
	}

	server = ServerBody{}
	json.NewDecoder(res.Body).Decode(&server)

	return dataSourceServerReadResponseBody(d, m, &server)
}

// resourceServerRead reads information about an existing server.
func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	req, reqErr := getClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s", d.Id()), new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	} else if res.StatusCode != 200 {
		if res.StatusCode == 404 {
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Failed to read the server information - Reason: The API responded with HTTP %s", res.Status)
	}

	server := ServerBody{}
	json.NewDecoder(res.Body).Decode(&server)

	return dataSourceServerReadResponseBody(d, m, &server)
}

// resourceServerUpdate updates an existing server.
func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	body := ServerUpdateBody{
		Hostname: d.Get(ResourceServerHostname).(string),
		Label:    d.Get(ResourceServerLabel).(string),
	}

	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(body)

	res, resErr := doClientRequest(&clientSettings, "PUT", fmt.Sprintf("cloudservers/%s", d.Id()), reqBody, []int{200}, 60, 10)

	if resErr != nil {
		return resErr
	}

	server := ServerBody{}
	json.NewDecoder(res.Body).Decode(&server)

	return dataSourceServerReadResponseBody(d, m, &server)
}

// resourceServerDelete deletes an existing server.
func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	_, err := doClientRequest(&clientSettings, "DELETE", fmt.Sprintf("cloudservers/%s", d.Id()), new(bytes.Buffer), []int{200, 404}, 60, 10)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
