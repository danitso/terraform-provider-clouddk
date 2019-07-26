package main

// ClientSettings describes the client settings.
type ClientSettings struct {
	Endpoint string
	Key      string
}

// DiskBody describes a disk object.
type DiskBody struct {
	Identifier string `json:"identifier"`
	Label      string `json:"label"`
	Size       string `json:"size"`
	Primary    uint8  `json:"primary"`
}

// DiskListBody describes a disk list.
type DiskListBody []DiskBody

// FirewallRuleBody describes a firewall rule object.
type FirewallRuleBody struct {
	Identifier string `json:"identifier"`
	Position   string `json:"position"`
	Command    string `json:"command"`
	Protocol   string `json:"protocol"`
	Address    string `json:"address"`
	Bits       string `json:"bits"`
	Port       string `json:"port"`
}

// FirewallRuleListBody describes a firewall rule list.
type FirewallRuleListBody []FirewallRuleBody

// IPAddressBody describes an IP address object.
type IPAddressBody struct {
	Address                    string `json:"address"`
	Network                    string `json:"network"`
	Netmask                    string `json:"netmask"`
	Gateway                    string `json:"gateway"`
	NetworkInterfaceIdentifier string `json:"network_interface_identifier"`
}

// IPAddressListBody describes an IP address list.
type IPAddressListBody []IPAddressBody

// LocationBody describes a datacenter location object.
type LocationBody struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

// LocationListBody describes a datacenter location list.
type LocationListBody []LocationBody

// NetworkInterfaceBody describes a network interface object.
type NetworkInterfaceBody struct {
	Identifier          string               `json:"identifier"`
	Label               string               `json:"label"`
	RateLimit           uint32               `json:"rate_limit"`
	DefaultFirewallRuke string               `json:"default_firewall_rule"`
	Primary             string               `json:"primary"`
	IPAddresses         IPAddressListBody    `json:"ipAddresses"`
	FirewallRules       FirewallRuleListBody `json:"firewallRules"`
}

// NetworkInterfaceListBody describes a network interface list.
type NetworkInterfaceListBody []NetworkInterfaceBody

// PackageBody describes a server package object.
type PackageBody struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

// PackageeListBody describes a server package list.
type PackageeListBody []PackageBody

// ServerBody describes a server object.
type ServerBody struct {
	Identifier        string                   `json:"identifier"`
	Hostname          string                   `json:"hostname"`
	Label             string                   `json:"label"`
	CPUs              string                   `json:"cpus"`
	Memory            string                   `json:"memory"`
	Booted            uint8                    `json:"booted"`
	Disks             DiskListBody             `json:"disks"`
	NetworkInterfaces NetworkInterfaceListBody `json:"networkInterfaces"`
	Template          TemplateBody             `json:"template"`
	Location          LocationBody             `json:"location"`
	Package           PackageBody              `json:"package"`
}

// ServerListBody describes a server list.
type ServerListBody []ServerBody

// TemplateBody describes a datacenter location object.
type TemplateBody struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

// TemplateListBody describes a datacenter location list.
type TemplateListBody []TemplateBody
