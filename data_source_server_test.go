package main

import (
	"testing"
)

// TestDataSourceServerInstantiation() tests whether the dataSourceServer instance can be instantiated.
func TestDataSourceServerInstantiation(t *testing.T) {
	s := dataSourceServer()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceServer")
	}
}

// TestDataSourceServerSchema() tests the dataSourceServer schema.
func TestDataSourceServerSchema(t *testing.T) {
	s := dataSourceServer()

	if s.Schema[DataSourceServerIdKey] == nil {
		t.Fatalf("Error in dataSourceServer.Schema: Missing argument \"%s\"", DataSourceServerIdKey)
	}

	if s.Schema[DataSourceServerIdKey].Required != true {
		t.Fatalf("Error in dataSourceServer.Schema: Argument \"%s\" is not required", DataSourceServerIdKey)
	}

	attributeKeys := []string{
		DataSourceServerBootedKey,
		DataSourceServerCPUsKey,
		DataSourceServerDiskIdsKey,
		DataSourceServerDiskLabelsKey,
		DataSourceServerDiskPrimaryKey,
		DataSourceServerDiskSizesKey,
		DataSourceServerHostnameKey,
		DataSourceServerLabelKey,
		DataSourceServerMemoryKey,
		DataSourceServerNetworkInterfaceAddressesKey,
		DataSourceServerNetworkInterfaceDefaultFirewallRulesKey,
		DataSourceServerNetworkInterfaceFirewallRulesAddressesKey,
		DataSourceServerNetworkInterfaceFirewallRulesCommandsKey,
		DataSourceServerNetworkInterfaceFirewallRulesIdsKey,
		DataSourceServerNetworkInterfaceFirewallRulesPortsKey,
		DataSourceServerNetworkInterfaceFirewallRulesProtocolsKey,
		DataSourceServerNetworkInterfaceGatewaysKey,
		DataSourceServerNetworkInterfaceIdsKey,
		DataSourceServerNetworkInterfaceLabelsKey,
		DataSourceServerNetworkInterfaceNetmasksKey,
		DataSourceServerNetworkInterfaceNetworksKey,
		DataSourceServerNetworkInterfacePrimaryKey,
		DataSourceServerNetworkInterfaceRateLimitsKey,
		DataSourceServerLocationIdKey,
		DataSourceServerLocationNameKey,
		DataSourceServerPackageIdKey,
		DataSourceServerPackageNameKey,
		DataSourceServerTemplateIdKey,
		DataSourceServerTemplateNameKey,
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
