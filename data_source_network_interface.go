package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceNetworkInterfaceAddressesKey = "addresses"
const DataSourceNetworkInterfaceDefaultFirewallRuleKey = "default_firewall_rule"
const DataSourceNetworkInterfaceFirewallRulesAddressesKey = "firewall_rules_addresses"
const DataSourceNetworkInterfaceFirewallRulesCommandsKey = "firewall_rules_commands"
const DataSourceNetworkInterfaceFirewallRulesIdsKey = "firewall_rules_ids"
const DataSourceNetworkInterfaceFirewallRulesPortsKey = "firewall_rules_ports"
const DataSourceNetworkInterfaceFirewallRulesProtocolsKey = "firewall_rules_protocols"
const DataSourceNetworkInterfaceGatewaysKey = "gateways"
const DataSourceNetworkInterfaceIdKey = "id"
const DataSourceNetworkInterfaceLabelKey = "label"
const DataSourceNetworkInterfaceNetmasksKey = "netmasks"
const DataSourceNetworkInterfaceNetworksKey = "networks"
const DataSourceNetworkInterfacePrimaryKey = "primary"
const DataSourceNetworkInterfaceRateLimitKey = "rate_limit"
const DataSourceNetworkInterfaceServerIdKey = "server_id"

// dataSourceNetworkInterface() retrieves information about a server.
func dataSourceNetworkInterface() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceNetworkInterfaceAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The IP addresses assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceNetworkInterfaceDefaultFirewallRuleKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default firewall rule for the network interface",
			},
			DataSourceNetworkInterfaceFirewallRulesAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The CIDR blocks for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceNetworkInterfaceFirewallRulesCommandsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The commands for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceNetworkInterfaceFirewallRulesIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The identifiers for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceNetworkInterfaceFirewallRulesPortsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The ports of the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceNetworkInterfaceFirewallRulesProtocolsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The protocols for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceNetworkInterfaceGatewaysKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The gateways assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceNetworkInterfaceIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The network interface identifier",
				ForceNew:    true,
			},
			DataSourceNetworkInterfaceLabelKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The network interface label",
			},
			DataSourceNetworkInterfaceNetmasksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The netmasks assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceNetworkInterfaceNetworksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The networks assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceNetworkInterfacePrimaryKey: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the network interface is the primary interface",
			},
			DataSourceNetworkInterfaceRateLimitKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The rate limit for the network interface",
			},
			DataSourceNetworkInterfaceServerIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
		},

		Read: dataSourceNetworkInterfaceRead,
	}
}

// dataSourceNetworkInterfaceRead() reads information about a server.
func dataSourceNetworkInterfaceRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	networkInterfaceId := d.Get(DataSourceNetworkInterfaceIdKey).(string)
	serverId := d.Get(DataSourceNetworkInterfaceServerIdKey).(string)

	req, reqErr := getClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/network-interfaces/%s", serverId, networkInterfaceId), new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the network interface - Reason: The API responded with HTTP %s", res.Status)
	}

	networkInterface := NetworkInterfaceBody{}
	json.NewDecoder(res.Body).Decode(&networkInterface)

	addresses := make([]interface{}, len(networkInterface.IPAddresses))
	gateways := make([]interface{}, len(networkInterface.IPAddresses))
	netmasks := make([]interface{}, len(networkInterface.IPAddresses))
	networks := make([]interface{}, len(networkInterface.IPAddresses))

	for i, v := range networkInterface.IPAddresses {
		addresses[i] = v.Address
		gateways[i] = v.Gateway
		netmasks[i] = v.Netmask
		networks[i] = v.Network
	}

	firewallRulesAddresses := make([]interface{}, len(networkInterface.FirewallRules))
	firewallRulesCommands := make([]interface{}, len(networkInterface.FirewallRules))
	firewallRulesIds := make([]interface{}, len(networkInterface.FirewallRules))
	firewallRulesPorts := make([]interface{}, len(networkInterface.FirewallRules))
	firewallRulesProtocols := make([]interface{}, len(networkInterface.FirewallRules))

	for _, v := range networkInterface.FirewallRules {
		firewallRulesAddresses[v.Position-1] = fmt.Sprintf("%s/%d", v.Address, v.Bits)
		firewallRulesCommands[v.Position-1] = v.Command
		firewallRulesIds[v.Position-1] = v.Identifier
		firewallRulesPorts[v.Position-1] = v.Port
		firewallRulesProtocols[v.Position-1] = v.Protocol
	}

	d.SetId(networkInterfaceId)

	d.Set(DataSourceNetworkInterfaceAddressesKey, addresses)

	d.Set(DataSourceNetworkInterfaceFirewallRulesAddressesKey, firewallRulesAddresses)
	d.Set(DataSourceNetworkInterfaceFirewallRulesCommandsKey, firewallRulesCommands)
	d.Set(DataSourceNetworkInterfaceFirewallRulesIdsKey, firewallRulesIds)
	d.Set(DataSourceNetworkInterfaceFirewallRulesPortsKey, firewallRulesPorts)
	d.Set(DataSourceNetworkInterfaceFirewallRulesProtocolsKey, firewallRulesProtocols)

	d.Set(DataSourceNetworkInterfaceGatewaysKey, gateways)
	d.Set(DataSourceNetworkInterfaceDefaultFirewallRuleKey, networkInterface.DefaultFirewallRule)
	d.Set(DataSourceNetworkInterfaceLabelKey, networkInterface.Label)
	d.Set(DataSourceNetworkInterfaceNetmasksKey, netmasks)
	d.Set(DataSourceNetworkInterfaceNetworksKey, networks)
	d.Set(DataSourceNetworkInterfacePrimaryKey, (networkInterface.Primary == 1))
	d.Set(DataSourceNetworkInterfaceRateLimitKey, networkInterface.RateLimit)

	return nil
}
