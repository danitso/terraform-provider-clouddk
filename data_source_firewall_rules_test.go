package main

import (
	"testing"
)

// TestDataSourceFirewallRulesInstantiation tests whether the dataSourceFirewallRules instance can be instantiated.
func TestDataSourceFirewallRulesInstantiation(t *testing.T) {
	s := dataSourceFirewallRules()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceFirewallRules")
	}
}

// TestDataSourceFirewallRulesSchema tests the dataSourceFirewallRules schema.
func TestDataSourceFirewallRulesSchema(t *testing.T) {
	s := dataSourceFirewallRules()

	idKeys := []string{
		dataSourceFirewallRulesIDKey,
		dataSourceFirewallRulesServerIDKey,
	}

	for _, v := range idKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceFirewallRules.Schema: Missing argument \"%s\"", v)
		}

		if s.Schema[v].Required != true {
			t.Fatalf("Error in dataSourceFirewallRules.Schema: Argument \"%s\" is not required", v)
		}
	}

	attributeKeys := []string{
		dataSourceFirewallRulesAddressesKey,
		dataSourceFirewallRulesCommandsKey,
		dataSourceFirewallRulesIdsKey,
		dataSourceFirewallRulesPortsKey,
		dataSourceFirewallRulesProtocolsKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceFirewallRules.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in dataSourceFirewallRules.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}
