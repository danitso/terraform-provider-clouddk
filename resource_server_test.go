/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"testing"
)

// TestResourceServerInstantiation tests whether the resourceServer instance can be instantiated.
func TestResourceServerInstantiation(t *testing.T) {
	s := resourceServer()

	if s == nil {
		t.Fatalf("Cannot instantiate resourceServer")
	}
}

// TestResourceServerSchema tests the resourceServer schema.
func TestResourceServerSchema(t *testing.T) {
	s := resourceServer()

	requiredKeys := []string{
		resourceServerHostnameKey,
		resourceServerLabelKey,
		resourceServerLocationIDKey,
		resourceServerPackageIDKey,
		resourceServerRootPasswordKey,
		resourceServerTemplateIDKey,
	}

	for _, v := range requiredKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in resourceServer.Schema: Missing argument \"%s\"", v)
		}

		if s.Schema[v].Required != true {
			t.Fatalf("Error in resourceServer.Schema: Argument \"%s\" is not required", v)
		}
	}

	optionalKeys := []string{
		resourceServerPrimaryNetworkInterfaceDefaultFirewallRuleKey,
		resourceServerPrimaryNetworkInterfaceLabelKey,
	}

	for _, v := range optionalKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in resourceServer.Schema: Missing argument \"%s\"", v)
		}

		if s.Schema[v].Optional != true {
			t.Fatalf("Error in resourceServer.Schema: Argument \"%s\" is not optional", v)
		}
	}

	attributeKeys := []string{
		dataSourceServerBootedKey,
		dataSourceServerCPUsKey,
		dataSourceServerDiskIdsKey,
		dataSourceServerDiskLabelsKey,
		dataSourceServerDiskPrimaryKey,
		dataSourceServerDiskSizesKey,
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
		dataSourceServerLocationNameKey,
		dataSourceServerPackageNameKey,
		dataSourceServerTemplateNameKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in resourceServer.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in resourceServer.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}
