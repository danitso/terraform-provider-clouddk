/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

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
	dataSourceNetworkInterfaceAddressesKey              = "addresses"
	dataSourceNetworkInterfaceDefaultFirewallRuleKey    = "default_firewall_rule"
	dataSourceNetworkInterfaceFirewallRulesAddressesKey = "firewall_rules_addresses"
	dataSourceNetworkInterfaceFirewallRulesCommandsKey  = "firewall_rules_commands"
	dataSourceNetworkInterfaceFirewallRulesIdsKey       = "firewall_rules_ids"
	dataSourceNetworkInterfaceFirewallRulesPortsKey     = "firewall_rules_ports"
	dataSourceNetworkInterfaceFirewallRulesProtocolsKey = "firewall_rules_protocols"
	dataSourceNetworkInterfaceGatewaysKey               = "gateways"
	dataSourceNetworkInterfaceIDKey                     = "id"
	dataSourceNetworkInterfaceLabelKey                  = "label"
	dataSourceNetworkInterfaceNetmasksKey               = "netmasks"
	dataSourceNetworkInterfaceNetworksKey               = "networks"
	dataSourceNetworkInterfacePrimaryKey                = "primary"
	dataSourceNetworkInterfaceRateLimitKey              = "rate_limit"
	dataSourceNetworkInterfaceServerIDKey               = "server_id"
)

// dataSourceNetworkInterface retrieves information about a server.
func dataSourceNetworkInterface() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourceNetworkInterfaceAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The IP addresses assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceNetworkInterfaceDefaultFirewallRuleKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default firewall rule for the network interface",
			},
			dataSourceNetworkInterfaceFirewallRulesAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The CIDR blocks for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceNetworkInterfaceFirewallRulesCommandsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The commands for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceNetworkInterfaceFirewallRulesIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The identifiers for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceNetworkInterfaceFirewallRulesPortsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The ports of the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceNetworkInterfaceFirewallRulesProtocolsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The protocols for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceNetworkInterfaceGatewaysKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The gateways assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceNetworkInterfaceIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The network interface identifier",
				ForceNew:    true,
			},
			dataSourceNetworkInterfaceLabelKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The network interface label",
			},
			dataSourceNetworkInterfaceNetmasksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The netmasks assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceNetworkInterfaceNetworksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The networks assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceNetworkInterfacePrimaryKey: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the network interface is the primary interface",
			},
			dataSourceNetworkInterfaceRateLimitKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The rate limit for the network interface",
			},
			dataSourceNetworkInterfaceServerIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
		},

		Read: dataSourceNetworkInterfaceRead,
	}
}

// dataSourceNetworkInterfaceRead reads information about a server.
func dataSourceNetworkInterfaceRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	networkInterfaceID := d.Get(dataSourceNetworkInterfaceIDKey).(string)
	serverID := d.Get(dataSourceNetworkInterfaceServerIDKey).(string)

	req, err := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/network-interfaces/%s", serverID, networkInterfaceID), new(bytes.Buffer))

	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the network interface - Reason: The API responded with HTTP %s", res.Status)
	}

	networkInterface := clouddk.NetworkInterfaceBody{}
	err = json.NewDecoder(res.Body).Decode(&networkInterface)

	if err != nil {
		return err
	}

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

	d.SetId(networkInterfaceID)

	d.Set(dataSourceNetworkInterfaceAddressesKey, addresses)

	d.Set(dataSourceNetworkInterfaceFirewallRulesAddressesKey, firewallRulesAddresses)
	d.Set(dataSourceNetworkInterfaceFirewallRulesCommandsKey, firewallRulesCommands)
	d.Set(dataSourceNetworkInterfaceFirewallRulesIdsKey, firewallRulesIds)
	d.Set(dataSourceNetworkInterfaceFirewallRulesPortsKey, firewallRulesPorts)
	d.Set(dataSourceNetworkInterfaceFirewallRulesProtocolsKey, firewallRulesProtocols)

	d.Set(dataSourceNetworkInterfaceGatewaysKey, gateways)
	d.Set(dataSourceNetworkInterfaceDefaultFirewallRuleKey, networkInterface.DefaultFirewallRule)
	d.Set(dataSourceNetworkInterfaceLabelKey, networkInterface.Label)
	d.Set(dataSourceNetworkInterfaceNetmasksKey, netmasks)
	d.Set(dataSourceNetworkInterfaceNetworksKey, networks)
	d.Set(dataSourceNetworkInterfacePrimaryKey, networkInterface.Primary)
	d.Set(dataSourceNetworkInterfaceRateLimitKey, networkInterface.RateLimit)

	return nil
}
