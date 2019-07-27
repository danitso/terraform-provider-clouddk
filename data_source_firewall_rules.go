package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceFirewallRulesAddressesKey = "addresses"
const DataSourceFirewallRulesCommandsKey = "commands"
const DataSourceFirewallRulesIdKey = "id"
const DataSourceFirewallRulesIdsKey = "ids"
const DataSourceFirewallRulesPortsKey = "ports"
const DataSourceFirewallRulesProtocolsKey = "protocols"
const DataSourceFirewallRulesServerIdKey = "server_id"

// dataSourceFirewallRules() retrieves information about firewall rules for a network interface.
func dataSourceFirewallRules() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceFirewallRulesAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The CIDR blocks for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceFirewallRulesCommandsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The commands for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceFirewallRulesIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The network interface identifier",
				ForceNew:    true,
			},
			DataSourceFirewallRulesIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The identifiers for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceFirewallRulesPortsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The ports of the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceFirewallRulesProtocolsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The protocols for the firewall rules assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceFirewallRulesServerIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
		},

		Read: dataSourceFirewallRulesRead,
	}
}

// dataSourceFirewallRulesRead() reads information about firewall rules for a network interface.
func dataSourceFirewallRulesRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	networkInterfaceId := d.Get(DataSourceFirewallRulesIdKey).(string)
	serverId := d.Get(DataSourceFirewallRulesServerIdKey).(string)

	req, reqErr := getClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/network-interfaces/%s/firewall-rules", serverId, networkInterfaceId), new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the firewall rules - Reason: The API responded with HTTP %s", res.Status)
	}

	firewallRules := FirewallRuleListBody{}
	json.NewDecoder(res.Body).Decode(&firewallRules)

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

	d.SetId(networkInterfaceId)

	d.Set(DataSourceFirewallRulesAddressesKey, firewallRulesAddresses)
	d.Set(DataSourceFirewallRulesCommandsKey, firewallRulesCommands)
	d.Set(DataSourceFirewallRulesIdsKey, firewallRulesIds)
	d.Set(DataSourceFirewallRulesPortsKey, firewallRulesPorts)
	d.Set(DataSourceFirewallRulesProtocolsKey, firewallRulesProtocols)

	return nil
}
