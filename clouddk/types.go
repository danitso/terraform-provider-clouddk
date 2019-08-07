package clouddk

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

// CustomBool allows a JSON boolean value to also be an integer
type CustomBool bool

// UnmarshalJSON converts a JSON value to a boolean.
func (r *CustomBool) UnmarshalJSON(b []byte) error {
	s := string(b)
	*r = CustomBool(s == "1" || s == "true")

	return nil
}

// MarshalJSON converts a boolean to a JSON value.
func (r CustomBool) MarshalJSON() ([]byte, error) {
	buffer := new(bytes.Buffer)

	if r {
		buffer.WriteString("1")
	} else {
		buffer.WriteString("0")
	}

	return buffer.Bytes(), nil
}

// CustomInt allows a JSON integer value to also be a string
type CustomInt int

// UnmarshalJSON converts a JSON value to an integer.
func (r *CustomInt) UnmarshalJSON(b []byte) error {
	s := string(b)

	if strings.Contains(s, "\"") {
		var v interface{}

		err := json.Unmarshal(b, &v)

		if err != nil {
			return err
		}

		s = v.(string)
	}

	i, err := strconv.Atoi(s)

	if err != nil {
		return err
	}

	*r = CustomInt(i)

	return nil
}

// ClientSettings describes the client settings.
type ClientSettings struct {
	Endpoint string
	Key      string
}

// DiskBody describes a disk object.
type DiskBody struct {
	Identifier string     `json:"identifier"`
	Label      string     `json:"label"`
	Size       CustomInt  `json:"size"`
	Primary    CustomBool `json:"primary"`
}

// DiskCreateBody describes a disk creation object.
type DiskCreateBody struct {
	Label string    `json:"label"`
	Size  CustomInt `json:"size"`
}

// DiskListBody describes a disk list.
type DiskListBody []DiskBody

// ErrorBody describes an error object.
type ErrorBody struct {
	Message string    `json:"message"`
	Status  CustomInt `json:"status"`
}

// FirewallRuleBody describes a firewall rule object.
type FirewallRuleBody struct {
	Identifier string    `json:"identifier"`
	Position   CustomInt `json:"position"`
	Command    string    `json:"command"`
	Protocol   string    `json:"protocol"`
	Address    string    `json:"address"`
	Bits       CustomInt `json:"bits"`
	Port       string    `json:"port"`
}

// FirewallRuleCreateBody describes a firewall rule creation object.
type FirewallRuleCreateBody struct {
	Command  string    `json:"command"`
	Protocol string    `json:"protocol"`
	Address  string    `json:"address"`
	Bits     CustomInt `json:"bits"`
	Port     string    `json:"port"`
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
	Identifier       CustomInt `json:"id"`
	Action           string    `json:"action"`
	Status           string    `json:"status"`
	TargetType       string    `json:"target_type"`
	TargetIdentifier CustomInt `json:"target_id"`
	CreatedAt        string    `json:"created_at"`
}

// LogsListBody describes a logs list.
type LogsListBody []LogsBody

// NetworkInterfaceBody describes a network interface object.
type NetworkInterfaceBody struct {
	Identifier          string               `json:"identifier"`
	Label               string               `json:"label"`
	RateLimit           CustomInt            `json:"rate_limit"`
	DefaultFirewallRule string               `json:"default_firewall_rule"`
	Primary             CustomBool           `json:"primary"`
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
	CPUs              CustomInt                `json:"cpus"`
	Memory            CustomInt                `json:"memory"`
	Booted            CustomBool               `json:"booted"`
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

// ServerUpgradeBody describes a server upgrade object.
type ServerUpgradeBody struct {
	Package     string     `json:"package"`
	UpgradeDisk CustomBool `json:"upgradeDisk"`
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
