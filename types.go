package main

// Bool allows a JSON boolean value to also be an integer
type Bool bool

func (bit *Bool) UnmarshalJSON(b []byte) error {
	txt := string(b)
	*bit = Bool(txt == "1" || txt == "true")

	return nil
}

// ClientSettings describes the client settings.
type ClientSettings struct {
	Endpoint string
	Key      string
}

// DiskBody describes a disk object.
type DiskBody struct {
	Identifier string `json:"identifier"`
	Label      string `json:"label"`
	Size       int    `json:"size"`
	Primary    int    `json:"primary"`
}

// DiskListBody describes a disk list.
type DiskListBody []DiskBody

// FirewallRuleBody describes a firewall rule object.
type FirewallRuleBody struct {
	Identifier string `json:"identifier"`
	Position   int    `json:"position"`
	Command    string `json:"command"`
	Protocol   string `json:"protocol"`
	Address    string `json:"address"`
	Bits       int    `json:"bits"`
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
	RateLimit           int                  `json:"rate_limit"`
	DefaultFirewallRule string               `json:"default_firewall_rule"`
	Primary             int                  `json:"primary"`
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
	CPUs              int                      `json:"cpus"`
	Memory            int                      `json:"memory"`
	Booted            Bool                     `json:"booted"`
	Disks             DiskListBody             `json:"disks"`
	NetworkInterfaces NetworkInterfaceListBody `json:"networkInterfaces"`
	Template          TemplateBody             `json:"template"`
	Location          LocationBody             `json:"location"`
	Package           PackageBody              `json:"package"`
}

// ServerCreateBody describes a server creation object.
type ServerCreateBody struct {
	Hostname            string `json:"hostname"`
	Label               string `json:"label"`
	InitialRootPassword string `json:"initialRootPassword"`
	Package             string `json:"package"`
	Template            string `json:"template"`
	Location            string `json:"location"`
}

// ServerListBody describes a server list.
type ServerListBody []ServerBody

// ServerUpdateBody describes a server update object.
type ServerUpdateBody struct {
	Hostname string `json:"hostname"`
	Label    string `json:"label"`
}

// TemplateBody describes a datacenter location object.
type TemplateBody struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

// TemplateListBody describes a datacenter location list.
type TemplateListBody []TemplateBody
