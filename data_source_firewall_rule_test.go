package main

import (
	"testing"
)

// TestDataSourceFirewallRuleInstantiation() tests whether the dataSourceFirewallRule instance can be instantiated.
func TestDataSourceFirewallRuleInstantiation(t *testing.T) {
	s := dataSourceFirewallRule()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceFirewallRule")
	}
}

// TestDataSourceFirewallRuleSchema() tests the dataSourceFirewallRule schema.
func TestDataSourceFirewallRuleSchema(t *testing.T) {
	s := dataSourceFirewallRule()

	idKeys := []string{
		DataSourceFirewallRuleIdKey,
		DataSourceFirewallRuleNetworkInterfaceIdKey,
		DataSourceFirewallRuleServerIdKey,
	}

	for _, v := range idKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceFirewallRule.Schema: Missing argument \"%s\"", v)
		}

		if s.Schema[v].Required != true {
			t.Fatalf("Error in dataSourceFirewallRule.Schema: Argument \"%s\" is not required", v)
		}
	}

	attributeKeys := []string{
		DataSourceFirewallRuleAddressKey,
		DataSourceFirewallRuleCommandKey,
		DataSourceFirewallRulePortKey,
		DataSourceFirewallRuleProtocolKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceFirewallRule.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in dataSourceFirewallRule.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}
