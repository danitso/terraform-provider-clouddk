package main

import (
	"testing"
)

// TestDataSourceNetworkInterfacesInstantiation() tests whether the dataSourceNetworkInterfaces instance can be instantiated.
func TestDataSourceNetworkInterfacesInstantiation(t *testing.T) {
	s := dataSourceNetworkInterfaces()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceNetworkInterfaces")
	}
}

// TestDataSourceNetworkInterfacesSchema() tests the dataSourceNetworkInterfaces schema.
func TestDataSourceNetworkInterfacesSchema(t *testing.T) {
	s := dataSourceNetworkInterfaces()

	if s.Schema[DataSourceNetworkInterfacesIdKey] == nil {
		t.Fatalf("Error in dataSourceNetworkInterfaces.Schema: Missing argument \"%s\"", DataSourceNetworkInterfacesIdKey)
	}

	if s.Schema[DataSourceNetworkInterfacesIdKey].Required != true {
		t.Fatalf("Error in dataSourceNetworkInterfaces.Schema: Argument \"%s\" is not required", DataSourceNetworkInterfacesIdKey)
	}

	attributeKeys := []string{
		DataSourceNetworkInterfacesAddressesKey,
		DataSourceNetworkInterfacesDefaultFirewallRulesKey,
		DataSourceNetworkInterfacesFirewallRulesAddressesKey,
		DataSourceNetworkInterfacesFirewallRulesCommandsKey,
		DataSourceNetworkInterfacesFirewallRulesIdsKey,
		DataSourceNetworkInterfacesFirewallRulesPortsKey,
		DataSourceNetworkInterfacesFirewallRulesProtocolsKey,
		DataSourceNetworkInterfacesGatewaysKey,
		DataSourceNetworkInterfacesIdsKey,
		DataSourceNetworkInterfacesLabelsKey,
		DataSourceNetworkInterfacesNetmasksKey,
		DataSourceNetworkInterfacesNetworksKey,
		DataSourceNetworkInterfacesPrimaryKey,
		DataSourceNetworkInterfacesRateLimitsKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceNetworkInterfaces.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in dataSourceNetworkInterfaces.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}
