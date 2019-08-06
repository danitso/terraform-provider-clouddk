package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	dataSourceNetworkInterfacesAddressesKey              = "addresses"
	dataSourceNetworkInterfacesDefaultFirewallRulesKey   = "default_firewall_rules"
	dataSourceNetworkInterfacesFirewallRulesAddressesKey = "firewall_rules_addresses"
	dataSourceNetworkInterfacesFirewallRulesCommandsKey  = "firewall_rules_commands"
	dataSourceNetworkInterfacesFirewallRulesIdsKey       = "firewall_rules_ids"
	dataSourceNetworkInterfacesFirewallRulesPortsKey     = "firewall_rules_ports"
	dataSourceNetworkInterfacesFirewallRulesProtocolsKey = "firewall_rules_protocols"
	dataSourceNetworkInterfacesGatewaysKey               = "gateways"
	dataSourceNetworkInterfacesIDKey                     = "id"
	dataSourceNetworkInterfacesIdsKey                    = "ids"
	dataSourceNetworkInterfacesLabelsKey                 = "labels"
	dataSourceNetworkInterfacesNetmasksKey               = "netmasks"
	dataSourceNetworkInterfacesNetworksKey               = "networks"
	dataSourceNetworkInterfacesPrimaryKey                = "primary"
	dataSourceNetworkInterfacesRateLimitsKey             = "rate_limits"
)

// dataSourceNetworkInterfaces retrieves information about a server.
func dataSourceNetworkInterfaces() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourceNetworkInterfacesAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The IP addresses assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceNetworkInterfacesDefaultFirewallRulesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The default firewall rules for the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceNetworkInterfacesFirewallRulesAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The CIDR blocks for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceNetworkInterfacesFirewallRulesCommandsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The commands for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceNetworkInterfacesFirewallRulesIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The identifiers for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceNetworkInterfacesFirewallRulesPortsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The ports of the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceNetworkInterfacesFirewallRulesProtocolsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The protocols for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceNetworkInterfacesGatewaysKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The gateways assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceNetworkInterfacesIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
			dataSourceNetworkInterfacesIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's network interface identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceNetworkInterfacesLabelsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's network interface labels",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceNetworkInterfacesNetmasksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The netmasks assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceNetworkInterfacesNetworksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The networks assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceNetworkInterfacesPrimaryKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whether a network interface is the primary interface",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			dataSourceNetworkInterfacesRateLimitsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The rate limits for the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},

		Read: dataSourceNetworkInterfacesRead,
	}
}

// dataSourceNetworkInterfacesRead reads information about a server.
func dataSourceNetworkInterfacesRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	id := d.Get(dataSourceNetworkInterfacesIDKey).(string)
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

	d.Set(dataSourceNetworkInterfacesAddressesKey, networkInterfaceAddresses)

	d.Set(dataSourceNetworkInterfacesFirewallRulesAddressesKey, networkInterfaceFirewallRuleAddresses)
	d.Set(dataSourceNetworkInterfacesFirewallRulesCommandsKey, networkInterfaceFirewallRuleCommands)
	d.Set(dataSourceNetworkInterfacesFirewallRulesIdsKey, networkInterfaceFirewallRuleIds)
	d.Set(dataSourceNetworkInterfacesFirewallRulesPortsKey, networkInterfaceFirewallRulePorts)
	d.Set(dataSourceNetworkInterfacesFirewallRulesProtocolsKey, networkInterfaceFirewallRuleProtocols)

	d.Set(dataSourceNetworkInterfacesGatewaysKey, networkInterfaceGateways)
	d.Set(dataSourceNetworkInterfacesDefaultFirewallRulesKey, networkInterfaceDefaultFirewallRules)
	d.Set(dataSourceNetworkInterfacesIdsKey, networkInterfaceIds)
	d.Set(dataSourceNetworkInterfacesLabelsKey, networkInterfaceLabels)
	d.Set(dataSourceNetworkInterfacesNetmasksKey, networkInterfaceNetmasks)
	d.Set(dataSourceNetworkInterfacesNetworksKey, networkInterfaceNetworks)
	d.Set(dataSourceNetworkInterfacesPrimaryKey, networkInterfacePrimary)
	d.Set(dataSourceNetworkInterfacesRateLimitsKey, networkInterfaceRateLimits)

	return nil
}
