package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceServerBootedKey = "booted"
const DataSourceServerCPUsKey = "cpus"
const DataSourceServerDiskIdsKey = "disk_ids"
const DataSourceServerDiskLabelsKey = "disk_labels"
const DataSourceServerDiskPrimaryKey = "disk_primary"
const DataSourceServerDiskSizesKey = "disk_sizes"
const DataSourceServerHostnameKey = "hostname"
const DataSourceServerIdKey = "id"
const DataSourceServerLabelKey = "label"
const DataSourceServerMemoryKey = "memory"
const DataSourceServerNetworkInterfaceAddressesKey = "network_interface_addresses"
const DataSourceServerNetworkInterfaceDefaultFirewallRulesKey = "network_interface_default_firewall_rules"
const DataSourceServerNetworkInterfaceFirewallRuleAddressesKey = "network_interface_firewall_rule_addresses"
const DataSourceServerNetworkInterfaceFirewallRuleBitsKey = "network_interface_firewall_rule_bits"
const DataSourceServerNetworkInterfaceFirewallRuleCommandsKey = "network_interface_firewall_rule_commands"
const DataSourceServerNetworkInterfaceFirewallRuleIdsKey = "network_interface_firewall_rule_ids"
const DataSourceServerNetworkInterfaceFirewallRulePortsKey = "network_interface_firewall_rule_ports"
const DataSourceServerNetworkInterfaceFirewallRulePositionsKey = "network_interface_firewall_rule_positions"
const DataSourceServerNetworkInterfaceFirewallRuleProtocolsKey = "network_interface_firewall_rule_protocols"
const DataSourceServerNetworkInterfaceGatewaysKey = "network_interface_gateways"
const DataSourceServerNetworkInterfaceIdsKey = "network_interface_ids"
const DataSourceServerNetworkInterfaceLabelsKey = "network_interface_labels"
const DataSourceServerNetworkInterfaceNetmasksKey = "network_interface_netmasks"
const DataSourceServerNetworkInterfaceNetworksKey = "network_interface_networks"
const DataSourceServerNetworkInterfacePrimaryKey = "network_interface_primary"
const DataSourceServerNetworkInterfaceRateLimitsKey = "network_interface_rate_limits"
const DataSourceServerLocationIdKey = "location_id"
const DataSourceServerLocationNameKey = "location_name"
const DataSourceServerPackageIdKey = "package_id"
const DataSourceServerPackageNameKey = "package_name"
const DataSourceServerTemplateIdKey = "template_id"
const DataSourceServerTemplateNameKey = "template_name"

// dataSourceServer() retrieves information about a server.
func dataSourceServer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
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
				Description: "Whether the disk is the primary disk",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			DataSourceServerDiskSizesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk sizes in gigabytes",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			DataSourceServerHostnameKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The server hostname",
			},
			DataSourceServerIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
			DataSourceServerLabelKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The server label",
			},
			DataSourceServerLocationIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The location identifier",
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
			DataSourceServerNetworkInterfaceFirewallRuleAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The addresses for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceFirewallRuleBitsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The bitmasks of the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeInt},
				},
			},
			DataSourceServerNetworkInterfaceFirewallRuleCommandsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The commands for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceFirewallRuleIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The identifiers for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceFirewallRulePortsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The ports of the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeInt},
				},
			},
			DataSourceServerNetworkInterfaceFirewallRulePositionsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The position of the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeInt},
				},
			},
			DataSourceServerNetworkInterfaceFirewallRuleProtocolsKey: &schema.Schema{
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
			DataSourceServerPackageIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The package identifier",
			},
			DataSourceServerPackageNameKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The package name",
			},
			DataSourceServerTemplateIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The template identifier",
			},
			DataSourceServerTemplateNameKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The template name",
			},
		},

		Read: dataSourceServerRead,
	}
}

// dataSourceServerRead() reads information about a server.
func dataSourceServerRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	id := d.Get(DataSourceServerIdKey).(string)
	req, reqErr := getClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s", id), new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	}

	server := ServerBody{}
	json.NewDecoder(res.Body).Decode(&server)

	diskIds := make([]interface{}, len(server.Disks))
	diskLabels := make([]interface{}, len(server.Disks))
	diskPrimary := make([]interface{}, len(server.Disks))
	diskSizes := make([]interface{}, len(server.Disks))

	for i, v := range server.Disks {
		diskIds[i] = v.Identifier
		diskLabels[i] = v.Label
		diskPrimary[i] = (v.Primary == 1)
		diskSizes[i] = v.Size
	}

	networkInterfaceAddresses := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceDefaultFirewallRules := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceFirewallRuleAddresses := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceFirewallRuleBits := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceFirewallRuleCommands := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceFirewallRuleIds := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceFirewallRulePorts := make([]interface{}, len(server.NetworkInterfaces))
	networkInterfaceFirewallRulePositions := make([]interface{}, len(server.NetworkInterfaces))
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

		firewallRuleAddresses := make([]interface{}, len(v.FirewallRules))
		firewallRuleBits := make([]interface{}, len(v.FirewallRules))
		firewallRuleCommands := make([]interface{}, len(v.FirewallRules))
		firewallRuleIds := make([]interface{}, len(v.FirewallRules))
		firewallRulePorts := make([]interface{}, len(v.FirewallRules))
		firewallRulePositions := make([]interface{}, len(v.FirewallRules))
		firewallRuleProtocols := make([]interface{}, len(v.FirewallRules))

		for ia, va := range v.FirewallRules {
			firewallRuleAddresses[ia] = va.Address
			firewallRuleBits[ia] = va.Bits
			firewallRuleCommands[ia] = va.Command
			firewallRuleIds[ia] = va.Identifier
			firewallRulePorts[ia] = va.Port
			firewallRulePositions[ia] = va.Position
			firewallRuleProtocols[ia] = va.Protocol
		}

		networkInterfaceAddresses[i] = addresses
		networkInterfaceDefaultFirewallRules[i] = v.DefaultFirewallRule

		networkInterfaceFirewallRuleAddresses[i] = firewallRuleAddresses
		networkInterfaceFirewallRuleBits[i] = firewallRuleBits
		networkInterfaceFirewallRuleCommands[i] = firewallRuleCommands
		networkInterfaceFirewallRuleIds[i] = firewallRuleIds
		networkInterfaceFirewallRulePorts[i] = firewallRulePorts
		networkInterfaceFirewallRulePositions[i] = firewallRulePositions
		networkInterfaceFirewallRuleProtocols[i] = firewallRuleProtocols

		networkInterfaceGateways[i] = gateways
		networkInterfaceIds[i] = v.Identifier
		networkInterfaceLabels[i] = v.Label
		networkInterfaceNetmasks[i] = netmasks
		networkInterfaceNetworks[i] = networks
		networkInterfacePrimary[i] = (v.Primary == 1)
		networkInterfaceRateLimits[i] = v.RateLimit
	}

	d.SetId(id)

	d.Set(DataSourceServerBootedKey, (server.Booted == 1))
	d.Set(DataSourceServerCPUsKey, server.CPUs)
	d.Set(DataSourceServerDiskIdsKey, diskIds)
	d.Set(DataSourceServerDiskLabelsKey, diskLabels)
	d.Set(DataSourceServerDiskPrimaryKey, diskPrimary)
	d.Set(DataSourceServerDiskSizesKey, diskSizes)
	d.Set(DataSourceServerHostnameKey, server.Hostname)
	d.Set(DataSourceServerLabelKey, server.Label)
	d.Set(DataSourceServerLocationIdKey, server.Location.Identifier)
	d.Set(DataSourceServerLocationNameKey, server.Location.Name)
	d.Set(DataSourceServerMemoryKey, server.Memory)
	d.Set(DataSourceServerNetworkInterfaceAddressesKey, networkInterfaceAddresses)

	d.Set(DataSourceServerNetworkInterfaceFirewallRuleAddressesKey, networkInterfaceFirewallRuleAddresses)
	d.Set(DataSourceServerNetworkInterfaceFirewallRuleBitsKey, networkInterfaceFirewallRuleBits)
	d.Set(DataSourceServerNetworkInterfaceFirewallRuleCommandsKey, networkInterfaceFirewallRuleCommands)
	d.Set(DataSourceServerNetworkInterfaceFirewallRuleIdsKey, networkInterfaceFirewallRuleIds)
	d.Set(DataSourceServerNetworkInterfaceFirewallRulePortsKey, networkInterfaceFirewallRulePorts)
	d.Set(DataSourceServerNetworkInterfaceFirewallRulePositionsKey, networkInterfaceFirewallRulePositions)
	d.Set(DataSourceServerNetworkInterfaceFirewallRuleProtocolsKey, networkInterfaceFirewallRuleProtocols)

	d.Set(DataSourceServerNetworkInterfaceGatewaysKey, networkInterfaceGateways)
	d.Set(DataSourceServerNetworkInterfaceDefaultFirewallRulesKey, networkInterfaceDefaultFirewallRules)
	d.Set(DataSourceServerNetworkInterfaceIdsKey, networkInterfaceIds)
	d.Set(DataSourceServerNetworkInterfaceLabelsKey, networkInterfaceLabels)
	d.Set(DataSourceServerNetworkInterfaceNetmasksKey, networkInterfaceNetmasks)
	d.Set(DataSourceServerNetworkInterfaceNetworksKey, networkInterfaceNetworks)
	d.Set(DataSourceServerNetworkInterfacePrimaryKey, networkInterfacePrimary)
	d.Set(DataSourceServerNetworkInterfaceRateLimitsKey, networkInterfaceRateLimits)

	d.Set(DataSourceServerPackageIdKey, server.Package.Identifier)
	d.Set(DataSourceServerPackageNameKey, server.Package.Name)
	d.Set(DataSourceServerTemplateIdKey, server.Template.Identifier)
	d.Set(DataSourceServerTemplateNameKey, server.Template.Name)

	return nil
}
