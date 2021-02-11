/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package clouddktf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	dataSourceFirewallRulesAddressesKey = "addresses"
	dataSourceFirewallRulesCommandsKey  = "commands"
	dataSourceFirewallRulesIDKey        = "id"
	dataSourceFirewallRulesIdsKey       = "ids"
	dataSourceFirewallRulesPortsKey     = "ports"
	dataSourceFirewallRulesProtocolsKey = "protocols"
	dataSourceFirewallRulesServerIDKey  = "server_id"
)

// dataSourceFirewallRules retrieves information about firewall rules for a network interface.
func dataSourceFirewallRules() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourceFirewallRulesAddressesKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The CIDR blocks for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceFirewallRulesCommandsKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The commands for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceFirewallRulesIDKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The network interface identifier",
				ForceNew:    true,
			},
			dataSourceFirewallRulesIdsKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The identifiers for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceFirewallRulesPortsKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The ports of the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceFirewallRulesProtocolsKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The protocols for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceFirewallRulesServerIDKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
		},

		Read: dataSourceFirewallRulesRead,
	}
}

// dataSourceFirewallRulesRead reads information about firewall rules for a network interface.
func dataSourceFirewallRulesRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	networkInterfaceID := d.Get(dataSourceFirewallRulesIDKey).(string)
	serverID := d.Get(dataSourceFirewallRulesServerIDKey).(string)

	req, err := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/network-interfaces/%s/firewall-rules", serverID, networkInterfaceID), new(bytes.Buffer))

	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the firewall rules - Reason: The API responded with HTTP %s", res.Status)
	}

	firewallRules := clouddk.FirewallRuleListBody{}
	err = json.NewDecoder(res.Body).Decode(&firewallRules)

	if err != nil {
		return err
	}

	firewallRulesAddresses := make([]interface{}, len(firewallRules))
	firewallRulesCommands := make([]interface{}, len(firewallRules))
	firewallRulesIds := make([]interface{}, len(firewallRules))
	firewallRulesPorts := make([]interface{}, len(firewallRules))
	firewallRulesProtocols := make([]interface{}, len(firewallRules))

	for _, v := range firewallRules {
		firewallRulesAddresses[v.Position-1] = fmt.Sprintf("%s/%d", v.Address, v.Bits)
		firewallRulesCommands[v.Position-1] = v.Command
		firewallRulesIds[v.Position-1] = v.Identifier
		firewallRulesPorts[v.Position-1] = v.Port
		firewallRulesProtocols[v.Position-1] = v.Protocol
	}

	d.SetId(networkInterfaceID)

	d.Set(dataSourceFirewallRulesAddressesKey, firewallRulesAddresses)
	d.Set(dataSourceFirewallRulesCommandsKey, firewallRulesCommands)
	d.Set(dataSourceFirewallRulesIdsKey, firewallRulesIds)
	d.Set(dataSourceFirewallRulesPortsKey, firewallRulesPorts)
	d.Set(dataSourceFirewallRulesProtocolsKey, firewallRulesProtocols)

	return nil
}
