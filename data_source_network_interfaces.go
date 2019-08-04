package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceNetworkInterfacesAddressesKey = "addresses"
const DataSourceNetworkInterfacesDefaultFirewallRulesKey = "default_firewall_rules"
const DataSourceNetworkInterfacesFirewallRulesAddressesKey = "firewall_rules_addresses"
const DataSourceNetworkInterfacesFirewallRulesCommandsKey = "firewall_rules_commands"
const DataSourceNetworkInterfacesFirewallRulesIdsKey = "firewall_rules_ids"
const DataSourceNetworkInterfacesFirewallRulesPortsKey = "firewall_rules_ports"
const DataSourceNetworkInterfacesFirewallRulesProtocolsKey = "firewall_rules_protocols"
const DataSourceNetworkInterfacesGatewaysKey = "gateways"
const DataSourceNetworkInterfacesIdKey = "id"
const DataSourceNetworkInterfacesIdsKey = "ids"
const DataSourceNetworkInterfacesLabelsKey = "labels"
const DataSourceNetworkInterfacesNetmasksKey = "netmasks"
const DataSourceNetworkInterfacesNetworksKey = "networks"
const DataSourceNetworkInterfacesPrimaryKey = "primary"
const DataSourceNetworkInterfacesRateLimitsKey = "rate_limits"

// dataSourceNetworkInterfaces() retrieves information about a server.
func dataSourceNetworkInterfaces() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceNetworkInterfacesAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The IP addresses assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceNetworkInterfacesDefaultFirewallRulesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The default firewall rules for the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceNetworkInterfacesFirewallRulesAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The CIDR blocks for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceNetworkInterfacesFirewallRulesCommandsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The commands for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceNetworkInterfacesFirewallRulesIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The identifiers for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceNetworkInterfacesFirewallRulesPortsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The ports of the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceNetworkInterfacesFirewallRulesProtocolsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The protocols for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceNetworkInterfacesGatewaysKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The gateways assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceNetworkInterfacesIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
			DataSourceNetworkInterfacesIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's network interface identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceNetworkInterfacesLabelsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's network interface labels",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceNetworkInterfacesNetmasksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The netmasks assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceNetworkInterfacesNetworksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The networks assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceNetworkInterfacesPrimaryKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whether a network interface is the primary interface",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			DataSourceNetworkInterfacesRateLimitsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The rate limits for the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},

		Read: dataSourceNetworkInterfacesRead,
	}
}

// dataSourceNetworkInterfacesRead() reads information about a server.
func dataSourceNetworkInterfacesRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	id := d.Get(DataSourceNetworkInterfacesIdKey).(string)
	req, reqErr := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/network-interfaces", id), new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the network interfaces - Reason: The API responded with HTTP %s", res.Status)
	}

	networkInterfaces := clouddk.NetworkInterfaceListBody{}
	json.NewDecoder(res.Body).Decode(&networkInterfaces)

	networkInterfaceAddresses := make([]interface{}, len(networkInterfaces))
	networkInterfaceDefaultFirewallRules := make([]interface{}, len(networkInterfaces))
	networkInterfaceFirewallRuleAddresses := make([]interface{}, len(networkInterfaces))
	networkInterfaceFirewallRuleCommands := make([]interface{}, len(networkInterfaces))
	networkInterfaceFirewallRuleIds := make([]interface{}, len(networkInterfaces))
	networkInterfaceFirewallRulePorts := make([]interface{}, len(networkInterfaces))
	networkInterfaceFirewallRuleProtocols := make([]interface{}, len(networkInterfaces))
	networkInterfaceGateways := make([]interface{}, len(networkInterfaces))
	networkInterfaceIds := make([]interface{}, len(networkInterfaces))
	networkInterfaceLabels := make([]interface{}, len(networkInterfaces))
	networkInterfaceNetmasks := make([]interface{}, len(networkInterfaces))
	networkInterfaceNetworks := make([]interface{}, len(networkInterfaces))
	networkInterfacePrimary := make([]interface{}, len(networkInterfaces))
	networkInterfaceRateLimits := make([]interface{}, len(networkInterfaces))

	for i, v := range networkInterfaces {
		addresses := make([]interface{}, len(v.IPAddresses))
		gateways := make([]interface{}, len(v.IPAddresses))
		netmasks := make([]interface{}, len(v.IPAddresses))
		networks := make([]interface{}, len(v.IPAddresses))

		for ia, va := range v.IPAddresses {
			addresses[ia] = va.Address
			gateways[ia] = va.Gateway
			netmasks[ia] = va.Netmask
			networks[ia] = va.Network
		}

		firewallRulesAddresses := make([]interface{}, len(v.FirewallRules))
		firewallRulesCommands := make([]interface{}, len(v.FirewallRules))
		firewallRulesIds := make([]interface{}, len(v.FirewallRules))
		firewallRulesPorts := make([]interface{}, len(v.FirewallRules))
		firewallRulesProtocols := make([]interface{}, len(v.FirewallRules))

		for _, va := range v.FirewallRules {
			firewallRulesAddresses[va.Position-1] = fmt.Sprintf("%s/%d", va.Address, va.Bits)
			firewallRulesCommands[va.Position-1] = va.Command
			firewallRulesIds[va.Position-1] = va.Identifier
			firewallRulesPorts[va.Position-1] = va.Port
			firewallRulesProtocols[va.Position-1] = va.Protocol
		}

		networkInterfaceAddresses[i] = addresses
		networkInterfaceDefaultFirewallRules[i] = v.DefaultFirewallRule

		networkInterfaceFirewallRuleAddresses[i] = firewallRulesAddresses
		networkInterfaceFirewallRuleCommands[i] = firewallRulesCommands
		networkInterfaceFirewallRuleIds[i] = firewallRulesIds
		networkInterfaceFirewallRulePorts[i] = firewallRulesPorts
		networkInterfaceFirewallRuleProtocols[i] = firewallRulesProtocols

		networkInterfaceGateways[i] = gateways
		networkInterfaceIds[i] = v.Identifier
		networkInterfaceLabels[i] = v.Label
		networkInterfaceNetmasks[i] = netmasks
		networkInterfaceNetworks[i] = networks
		networkInterfacePrimary[i] = v.Primary
		networkInterfaceRateLimits[i] = v.RateLimit
	}

	d.SetId(id)

	d.Set(DataSourceNetworkInterfacesAddressesKey, networkInterfaceAddresses)

	d.Set(DataSourceNetworkInterfacesFirewallRulesAddressesKey, networkInterfaceFirewallRuleAddresses)
	d.Set(DataSourceNetworkInterfacesFirewallRulesCommandsKey, networkInterfaceFirewallRuleCommands)
	d.Set(DataSourceNetworkInterfacesFirewallRulesIdsKey, networkInterfaceFirewallRuleIds)
	d.Set(DataSourceNetworkInterfacesFirewallRulesPortsKey, networkInterfaceFirewallRulePorts)
	d.Set(DataSourceNetworkInterfacesFirewallRulesProtocolsKey, networkInterfaceFirewallRuleProtocols)

	d.Set(DataSourceNetworkInterfacesGatewaysKey, networkInterfaceGateways)
	d.Set(DataSourceNetworkInterfacesDefaultFirewallRulesKey, networkInterfaceDefaultFirewallRules)
	d.Set(DataSourceNetworkInterfacesIdsKey, networkInterfaceIds)
	d.Set(DataSourceNetworkInterfacesLabelsKey, networkInterfaceLabels)
	d.Set(DataSourceNetworkInterfacesNetmasksKey, networkInterfaceNetmasks)
	d.Set(DataSourceNetworkInterfacesNetworksKey, networkInterfaceNetworks)
	d.Set(DataSourceNetworkInterfacesPrimaryKey, networkInterfacePrimary)
	d.Set(DataSourceNetworkInterfacesRateLimitsKey, networkInterfaceRateLimits)

	return nil
}
