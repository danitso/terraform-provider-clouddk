package clouddk

import "bytes"

// Bool allows a JSON boolean value to also be an integer
type Bool bool

func (this *Bool) UnmarshalJSON(b []byte) error {
	txt := string(b)
	*this = Bool(txt == "1" || txt == "true")

	return nil
}

func (this Bool) MarshalJSON() ([]byte, error) {
	buffer := new(bytes.Buffer)

	if this {
		buffer.WriteString("1")
	} else {
		buffer.WriteString("0")
	}

	return buffer.Bytes(), nil
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
	Primary    Bool   `json:"primary"`
}

// DiskCreateBody describes a disk creation object.
type DiskCreateBody struct {
	Label string `json:"label"`
	Size  int    `json:"size"`
}

// DiskListBody describes a disk list.
type DiskListBody []DiskBody

// ErrorBody describes an error object.
type ErrorBody struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

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

// FirewallRuleCreateBody describes a firewall rule creation object.
type FirewallRuleCreateBody struct {
	Command  string `json:"command"`
	Protocol string `json:"protocol"`
	Address  string `json:"address"`
	Bits     int    `json:"bits"`
	Port     string `json:"port"`
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

// LogsBody describes a logs object.
type LogsBody struct {
	Identifier       int    `json:"id"`
	Action           string `json:"action"`
	Status           string `json:"status"`
	TargetType       string `json:"target_type"`
	TargetIdentifier int    `json:"target_id"`
	CreatedAt        string `json:"created_at"`
}

// LogsListBody describes a logs list.
type LogsListBody []LogsBody

// NetworkInterfaceBody describes a network interface object.
type NetworkInterfaceBody struct {
	Identifier          string               `json:"identifier"`
	Label               string               `json:"label"`
	RateLimit           int                  `json:"rate_limit"`
	DefaultFirewallRule string               `json:"default_firewall_rule"`
	Primary             Bool                 `json:"primary"`
	IPAddresses         IPAddressListBody    `json:"ipAddresses"`
	FirewallRules       FirewallRuleListBody `json:"firewallRules"`
}

// NetworkInterfaceListBody describes a network interface list.
type NetworkInterfaceListBody []NetworkInterfaceBody

// NetworkInterfaceUpdateBody describes a network interface update object.
type NetworkInterfaceUpdateBody struct {
	Label               string `json:"label"`
	DefaultFirewallRule string `json:"default_firewall_rule"`
}

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
