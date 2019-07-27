package main

import (
	"testing"
)

// TestDataSourceIPAddressesInstantiation() tests whether the dataSourceIPAddresses instance can be instantiated.
func TestDataSourceIPAddressesInstantiation(t *testing.T) {
	s := dataSourceIPAddresses()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceIPAddresses")
	}
}

// TestDataSourceIPAddressesSchema() tests the dataSourceIPAddresses schema.
func TestDataSourceIPAddressesSchema(t *testing.T) {
	s := dataSourceIPAddresses()

	if s.Schema[DataSourceIPAddressesIdKey] == nil {
		t.Fatalf("Error in dataSourceIPAddresses.Schema: Missing argument \"%s\"", DataSourceIPAddressesIdKey)
	}

	if s.Schema[DataSourceIPAddressesIdKey].Required != true {
		t.Fatalf("Error in dataSourceIPAddresses.Schema: Argument \"%s\" is not required", DataSourceIPAddressesIdKey)
	}

	attributeKeys := []string{
		DataSourceIPAddressesAddressesKey,
		DataSourceIPAddressesGatewaysKey,
		DataSourceIPAddressesNetmasksKey,
		DataSourceIPAddressesNetworkInterfaceIdsKey,
		DataSourceIPAddressesNetworksKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceIPAddresses.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in dataSourceIPAddresses.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}
