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
	dataSourceServerBootedKey                                 = "booted"
	dataSourceServerCPUsKey                                   = "cpus"
	dataSourceServerDiskIdsKey                                = "disk_ids"
	dataSourceServerDiskLabelsKey                             = "disk_labels"
	dataSourceServerDiskPrimaryKey                            = "disk_primary"
	dataSourceServerDiskSizesKey                              = "disk_sizes"
	dataSourceServerHostnameKey                               = "hostname"
	dataSourceServerIDKey                                     = "id"
	dataSourceServerLabelKey                                  = "label"
	dataSourceServerMemoryKey                                 = "memory"
	dataSourceServerNetworkInterfaceAddressesKey              = "network_interface_addresses"
	dataSourceServerNetworkInterfaceDefaultFirewallRulesKey   = "network_interface_default_firewall_rules"
	dataSourceServerNetworkInterfaceFirewallRulesAddressesKey = "network_interface_firewall_rules_addresses"
	dataSourceServerNetworkInterfaceFirewallRulesCommandsKey  = "network_interface_firewall_rules_commands"
	dataSourceServerNetworkInterfaceFirewallRulesIdsKey       = "network_interface_firewall_rules_ids"
	dataSourceServerNetworkInterfaceFirewallRulesPortsKey     = "network_interface_firewall_rules_ports"
	dataSourceServerNetworkInterfaceFirewallRulesProtocolsKey = "network_interface_firewall_rules_protocols"
	dataSourceServerNetworkInterfaceGatewaysKey               = "network_interface_gateways"
	dataSourceServerNetworkInterfaceIdsKey                    = "network_interface_ids"
	dataSourceServerNetworkInterfaceLabelsKey                 = "network_interface_labels"
	dataSourceServerNetworkInterfaceNetmasksKey               = "network_interface_netmasks"
	dataSourceServerNetworkInterfaceNetworksKey               = "network_interface_networks"
	dataSourceServerNetworkInterfacePrimaryKey                = "network_interface_primary"
	dataSourceServerNetworkInterfaceRateLimitsKey             = "network_interface_rate_limits"
	dataSourceServerLocationIDKey                             = "location_id"
	dataSourceServerLocationNameKey                           = "location_name"
	dataSourceServerPackageIDKey                              = "package_id"
	dataSourceServerPackageNameKey                            = "package_name"
	dataSourceServerTemplateIDKey                             = "template_id"
	dataSourceServerTemplateNameKey                           = "template_name"
)

// dataSourceServer retrieves information about a server.
func dataSourceServer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourceServerBootedKey: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the server has been booted",
			},
			dataSourceServerCPUsKey: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The server's CPU count",
			},
			dataSourceServerDiskIdsKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServerDiskLabelsKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk labels",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServerDiskPrimaryKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whether a disk is the primary disk",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			dataSourceServerDiskSizesKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk sizes in gigabytes",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			dataSourceServerHostnameKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The server hostname",
			},
			dataSourceServerIDKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
			dataSourceServerLabelKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The server label",
			},
			dataSourceServerLocationIDKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The location identifier",
			},
			dataSourceServerLocationNameKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The location name",
			},
			dataSourceServerMemoryKey: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The server's memory allocation in megabytes",
			},
			dataSourceServerNetworkInterfaceAddressesKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The IP addresses assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceDefaultFirewallRulesKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The default firewall rules for the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServerNetworkInterfaceFirewallRulesAddressesKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The CIDR blocks for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceFirewallRulesCommandsKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The commands for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceFirewallRulesIdsKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The identifiers for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceFirewallRulesPortsKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The ports of the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceFirewallRulesProtocolsKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The protocols for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceGatewaysKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The gateways assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceIdsKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's network interface identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServerNetworkInterfaceLabelsKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's network interface labels",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServerNetworkInterfaceNetmasksKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The netmasks assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceNetworksKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The networks assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfacePrimaryKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whether a network interface is the primary interface",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			dataSourceServerNetworkInterfaceRateLimitsKey: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The rate limits for the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			dataSourceServerPackageIDKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The package identifier",
			},
			dataSourceServerPackageNameKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The package name",
			},
			dataSourceServerTemplateIDKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The template identifier",
			},
			dataSourceServerTemplateNameKey: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The template name",
			},
		},

		Read: dataSourceServerRead,
	}
}

// dataSourceServerRead reads information about a server.
func dataSourceServerRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)
	id := d.Id()

	if d.Get(dataSourceServerIDKey) != nil {
		id = d.Get(dataSourceServerIDKey).(string)
	}

	req, err := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s", id), new(bytes.Buffer))

	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the server - Reason: The API responded with HTTP %s", res.Status)
	}

	server := clouddk.ServerBody{}
	err = json.NewDecoder(res.Body).Decode(&server)

	if err != nil {
		return err
	}

	return dataSourceServerReadResponseBody(d, m, &server)
}

// dataSourceServerReadResponseBody() reads the response body for a server request.
func dataSourceServerReadResponseBody(d *schema.ResourceData, m interface{}, server *clouddk.ServerBody) error {
	diskIds := make([]interface{}, len(server.Disks))
	diskLabels := make([]interface{}, len(server.Disks))
	diskPrimary := make([]interface{}, len(server.Disks))
	diskSizes := make([]interface{}, len(server.Disks))

	for i, v := range server.Disks {
		diskIds[i] = v.Identifier
		diskLabels[i] = v.Label
		diskPrimary[i] = v.Primary
		diskSizes[i] = v.Size
	}

	networkInterfaceAddresses := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceDefaultFirewallRules := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceFirewallRuleAddresses := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceFirewallRuleCommands := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceFirewallRuleIds := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceFirewallRulePorts := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceFirewallRuleProtocols := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceGateways := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceIds := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceLabels := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceNetmasks := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceNetworks := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfacePrimary := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceRateLimits := make([]interface{}, len(server.NetworkInterfaces))

	for i, v := range server.NetworkInterfaces {
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

	d.SetId(server.Identifier)

	d.Set(dataSourceServerBootedKey, server.Booted)
	d.Set(dataSourceServerCPUsKey, server.CPUs)
	d.Set(dataSourceServerDiskIdsKey, diskIds)
	d.Set(dataSourceServerDiskLabelsKey, diskLabels)
	d.Set(dataSourceServerDiskPrimaryKey, diskPrimary)
	d.Set(dataSourceServerDiskSizesKey, diskSizes)
	d.Set(dataSourceServerHostnameKey, server.Hostname)
	d.Set(dataSourceServerLabelKey, server.Label)
	d.Set(dataSourceServerLocationIDKey, server.Location.Identifier)
	d.Set(dataSourceServerLocationNameKey, server.Location.Name)
	d.Set(dataSourceServerMemoryKey, server.Memory)
	d.Set(dataSourceServerNetworkInterfaceAddressesKey, networkInterfaceAddresses)

	d.Set(dataSourceServerNetworkInterfaceFirewallRulesAddressesKey, networkInterfaceFirewallRuleAddresses)
	d.Set(dataSourceServerNetworkInterfaceFirewallRulesCommandsKey, networkInterfaceFirewallRuleCommands)
	d.Set(dataSourceServerNetworkInterfaceFirewallRulesIdsKey, networkInterfaceFirewallRuleIds)
	d.Set(dataSourceServerNetworkInterfaceFirewallRulesPortsKey, networkInterfaceFirewallRulePorts)
	d.Set(dataSourceServerNetworkInterfaceFirewallRulesProtocolsKey, networkInterfaceFirewallRuleProtocols)

	d.Set(dataSourceServerNetworkInterfaceGatewaysKey, networkInterfaceGateways)
	d.Set(dataSourceServerNetworkInterfaceDefaultFirewallRulesKey, networkInterfaceDefaultFirewallRules)
	d.Set(dataSourceServerNetworkInterfaceIdsKey, networkInterfaceIds)
	d.Set(dataSourceServerNetworkInterfaceLabelsKey, networkInterfaceLabels)
	d.Set(dataSourceServerNetworkInterfaceNetmasksKey, networkInterfaceNetmasks)
	d.Set(dataSourceServerNetworkInterfaceNetworksKey, networkInterfaceNetworks)
	d.Set(dataSourceServerNetworkInterfacePrimaryKey, networkInterfacePrimary)
	d.Set(dataSourceServerNetworkInterfaceRateLimitsKey, networkInterfaceRateLimits)

	d.Set(dataSourceServerPackageIDKey, server.Package.Identifier)
	d.Set(dataSourceServerPackageNameKey, server.Package.Name)
	d.Set(dataSourceServerTemplateIDKey, server.Template.Identifier)
	d.Set(dataSourceServerTemplateNameKey, server.Template.Name)

	return nil
}
