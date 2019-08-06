package main

import (
	"testing"
)

// TestDataSourceServerInstantiation tests whether the dataSourceServer instance can be instantiated.
func TestDataSourceServerInstantiation(t *testing.T) {
	s := dataSourceServer()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceServer")
	}
}

// TestDataSourceServerSchema tests the dataSourceServer schema.
func TestDataSourceServerSchema(t *testing.T) {
	s := dataSourceServer()

	if s.Schema[dataSourceServerIDKey] == nil {
		t.Fatalf("Error in dataSourceServer.Schema: Missing argument \"%s\"", dataSourceServerIDKey)
	}

	if s.Schema[dataSourceServerIDKey].Required != true {
		t.Fatalf("Error in dataSourceServer.Schema: Argument \"%s\" is not required", dataSourceServerIDKey)
	}

	attributeKeys := []string{
		dataSourceServerBootedKey,
		dataSourceServerCPUsKey,
		dataSourceServerDiskIdsKey,
		dataSourceServerDiskLabelsKey,
		dataSourceServerDiskPrimaryKey,
		dataSourceServerDiskSizesKey,
		dataSourceServerHostnameKey,
		dataSourceServerLabelKey,
		dataSourceServerMemoryKey,
		dataSourceServerNetworkInterfaceAddressesKey,
		dataSourceServerNetworkInterfaceDefaultFirewallRulesKey,
		dataSourceServerNetworkInterfaceFirewallRulesAddressesKey,
		dataSourceServerNetworkInterfaceFirewallRulesCommandsKey,
		dataSourceServerNetworkInterfaceFirewallRulesIdsKey,
		dataSourceServerNetworkInterfaceFirewallRulesPortsKey,
		dataSourceServerNetworkInterfaceFirewallRulesProtocolsKey,
		dataSourceServerNetworkInterfaceGatewaysKey,
		dataSourceServerNetworkInterfaceIdsKey,
		dataSourceServerNetworkInterfaceLabelsKey,
		dataSourceServerNetworkInterfaceNetmasksKey,
		dataSourceServerNetworkInterfaceNetworksKey,
		dataSourceServerNetworkInterfacePrimaryKey,
		dataSourceServerNetworkInterfaceRateLimitsKey,
		dataSourceServerLocationIDKey,
		dataSourceServerLocationNameKey,
		dataSourceServerPackageIDKey,
		dataSourceServerPackageNameKey,
		dataSourceServerTemplateIDKey,
		dataSourceServerTemplateNameKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceServer.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in dataSourceServer.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}
