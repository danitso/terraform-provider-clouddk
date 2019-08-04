package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceFirewallRuleAddressKey = "address"
const DataSourceFirewallRuleCommandKey = "command"
const DataSourceFirewallRuleIdKey = "id"
const DataSourceFirewallRuleNetworkInterfaceIdKey = "network_interface_id"
const DataSourceFirewallRulePortKey = "port"
const DataSourceFirewallRuleProtocolKey = "protocol"
const DataSourceFirewallRuleServerIdKey = "server_id"

// dataSourceFirewallRule() retrieves information about a firewall rule for a network interface.
func dataSourceFirewallRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceFirewallRuleAddressKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CIDR block for the firewall rule",
			},
			DataSourceFirewallRuleCommandKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The command for the firewall rule",
			},
			DataSourceFirewallRuleIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The firewall rule identifier",
				ForceNew:    true,
			},
			DataSourceFirewallRuleNetworkInterfaceIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The network interface identifier",
				ForceNew:    true,
			},
			DataSourceFirewallRulePortKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The port for the firewall rule",
			},
			DataSourceFirewallRuleProtocolKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The protocol for the firewall rule",
			},
			DataSourceFirewallRuleServerIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
		},

		Read: dataSourceFirewallRuleRead,
	}
}

// dataSourceFirewallRuleRead() reads information about a firewall rule for a network interface.
func dataSourceFirewallRuleRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)
	firewallRuleId := d.Id()

	if d.Get(DataSourceFirewallRuleIdKey) != nil {
		firewallRuleId = d.Get(DataSourceFirewallRuleIdKey).(string)
	}

	networkInterfaceId := d.Get(DataSourceFirewallRuleNetworkInterfaceIdKey).(string)
	serverId := d.Get(DataSourceFirewallRuleServerIdKey).(string)

	req, reqErr := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/network-interfaces/%s/firewall-rules/%s", serverId, networkInterfaceId, firewallRuleId), new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the firewall rule - Reason: The API responded with HTTP %s", res.Status)
	}

	firewallRule := clouddk.FirewallRuleBody{}
	json.NewDecoder(res.Body).Decode(&firewallRule)

	return dataSourceFirewallRuleReadResponseBody(d, m, &firewallRule)
}

// dataSourceFirewallRuleReadResponseBody() reads information about a firewall rule for a network interface.
func dataSourceFirewallRuleReadResponseBody(d *schema.ResourceData, m interface{}, firewallRule *clouddk.FirewallRuleBody) error {
	d.SetId(firewallRule.Identifier)

	d.Set(DataSourceFirewallRuleAddressKey, fmt.Sprintf("%s/%d", firewallRule.Address, firewallRule.Bits))
	d.Set(DataSourceFirewallRuleCommandKey, firewallRule.Command)
	d.Set(DataSourceFirewallRulePortKey, firewallRule.Port)
	d.Set(DataSourceFirewallRuleProtocolKey, firewallRule.Protocol)

	return nil
}
