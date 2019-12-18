/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package clouddktf

import (
	"testing"
)

// TestDataSourceFirewallRuleInstantiation tests whether the dataSourceFirewallRule instance can be instantiated.
func TestDataSourceFirewallRuleInstantiation(t *testing.T) {
	s := dataSourceFirewallRule()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceFirewallRule")
	}
}

// TestDataSourceFirewallRuleSchema tests the dataSourceFirewallRule schema.
func TestDataSourceFirewallRuleSchema(t *testing.T) {
	s := dataSourceFirewallRule()

	idKeys := []string{
		dataSourceFirewallRuleIDKey,
		dataSourceFirewallRuleNetworkInterfaceIDKey,
		dataSourceFirewallRuleServerIDKey,
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
		dataSourceFirewallRuleAddressKey,
		dataSourceFirewallRuleCommandKey,
		dataSourceFirewallRulePortKey,
		dataSourceFirewallRuleProtocolKey,
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
