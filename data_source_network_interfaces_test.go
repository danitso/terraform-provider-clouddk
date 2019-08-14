/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"testing"
)

// TestDataSourceNetworkInterfacesInstantiation tests whether the dataSourceNetworkInterfaces instance can be instantiated.
func TestDataSourceNetworkInterfacesInstantiation(t *testing.T) {
	s := dataSourceNetworkInterfaces()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceNetworkInterfaces")
	}
}

// TestDataSourceNetworkInterfacesSchema tests the dataSourceNetworkInterfaces schema.
func TestDataSourceNetworkInterfacesSchema(t *testing.T) {
	s := dataSourceNetworkInterfaces()

	if s.Schema[dataSourceNetworkInterfacesIDKey] == nil {
		t.Fatalf("Error in dataSourceNetworkInterfaces.Schema: Missing argument \"%s\"", dataSourceNetworkInterfacesIDKey)
	}

	if s.Schema[dataSourceNetworkInterfacesIDKey].Required != true {
		t.Fatalf("Error in dataSourceNetworkInterfaces.Schema: Argument \"%s\" is not required", dataSourceNetworkInterfacesIDKey)
	}

	attributeKeys := []string{
		dataSourceNetworkInterfacesAddressesKey,
		dataSourceNetworkInterfacesDefaultFirewallRulesKey,
		dataSourceNetworkInterfacesFirewallRulesAddressesKey,
		dataSourceNetworkInterfacesFirewallRulesCommandsKey,
		dataSourceNetworkInterfacesFirewallRulesIdsKey,
		dataSourceNetworkInterfacesFirewallRulesPortsKey,
		dataSourceNetworkInterfacesFirewallRulesProtocolsKey,
		dataSourceNetworkInterfacesGatewaysKey,
		dataSourceNetworkInterfacesIdsKey,
		dataSourceNetworkInterfacesLabelsKey,
		dataSourceNetworkInterfacesNetmasksKey,
		dataSourceNetworkInterfacesNetworksKey,
		dataSourceNetworkInterfacesPrimaryKey,
		dataSourceNetworkInterfacesRateLimitsKey,
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
