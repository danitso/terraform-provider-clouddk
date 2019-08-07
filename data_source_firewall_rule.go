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
	dataSourceFirewallRuleAddressKey            = "address"
	dataSourceFirewallRuleCommandKey            = "command"
	dataSourceFirewallRuleIDKey                 = "id"
	dataSourceFirewallRuleNetworkInterfaceIDKey = "network_interface_id"
	dataSourceFirewallRulePortKey               = "port"
	dataSourceFirewallRuleProtocolKey           = "protocol"
	dataSourceFirewallRuleServerIDKey           = "server_id"
)

// dataSourceFirewallRule retrieves information about a firewall rule for a network interface.
func dataSourceFirewallRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourceFirewallRuleAddressKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CIDR block for the firewall rule",
			},
			dataSourceFirewallRuleCommandKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The command for the firewall rule",
			},
			dataSourceFirewallRuleIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The firewall rule identifier",
				ForceNew:    true,
			},
			dataSourceFirewallRuleNetworkInterfaceIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The network interface identifier",
				ForceNew:    true,
			},
			dataSourceFirewallRulePortKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The port for the firewall rule",
			},
			dataSourceFirewallRuleProtocolKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The protocol for the firewall rule",
			},
			dataSourceFirewallRuleServerIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
		},

		Read: dataSourceFirewallRuleRead,
	}
}

// dataSourceFirewallRuleRead reads information about a firewall rule for a network interface.
func dataSourceFirewallRuleRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)
	firewallRuleID := d.Id()

	if d.Get(dataSourceFirewallRuleIDKey) != nil {
		firewallRuleID = d.Get(dataSourceFirewallRuleIDKey).(string)
	}

	networkInterfaceID := d.Get(dataSourceFirewallRuleNetworkInterfaceIDKey).(string)
	serverID := d.Get(dataSourceFirewallRuleServerIDKey).(string)

	req, err := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/network-interfaces/%s/firewall-rules/%s", serverID, networkInterfaceID, firewallRuleID), new(bytes.Buffer))

	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the firewall rule - Reason: The API responded with HTTP %s", res.Status)
	}

	firewallRule := clouddk.FirewallRuleBody{}
	err = json.NewDecoder(res.Body).Decode(&firewallRule)

	if err != nil {
		return err
	}

	return dataSourceFirewallRuleReadResponseBody(d, m, &firewallRule)
}

// dataSourceFirewallRuleReadResponseBody reads information about a firewall rule for a network interface.
func dataSourceFirewallRuleReadResponseBody(d *schema.ResourceData, m interface{}, firewallRule *clouddk.FirewallRuleBody) error {
	d.SetId(firewallRule.Identifier)

	d.Set(dataSourceFirewallRuleAddressKey, fmt.Sprintf("%s/%d", firewallRule.Address, firewallRule.Bits))
	d.Set(dataSourceFirewallRuleCommandKey, firewallRule.Command)
	d.Set(dataSourceFirewallRulePortKey, firewallRule.Port)
	d.Set(dataSourceFirewallRuleProtocolKey, firewallRule.Protocol)

	return nil
}
