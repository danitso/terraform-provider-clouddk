package main

import (
	"testing"
)

// TestResourceServerInstantiation() tests whether the resourceServer instance can be instantiated.
func TestResourceServerInstantiation(t *testing.T) {
	s := resourceServer()

	if s == nil {
		t.Fatalf("Cannot instantiate resourceServer")
	}
}

// TestResourceServerSchema() tests the resourceServer schema.
func TestResourceServerSchema(t *testing.T) {
	s := resourceServer()

	requiredKeys := []string{
		ResourceServerHostnameKey,
		ResourceServerLabelKey,
		ResourceServerLocationIdKey,
		ResourceServerPackageIdKey,
		ResourceServerRootPasswordKey,
		ResourceServerTemplateIdKey,
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
		ResourceServerPrimaryNetworkInterfaceDefaultFirewallRuleKey,
		ResourceServerPrimaryNetworkInterfaceLabelKey,
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
		DataSourceServerBootedKey,
		DataSourceServerCPUsKey,
		DataSourceServerDiskIdsKey,
		DataSourceServerDiskLabelsKey,
		DataSourceServerDiskPrimaryKey,
		DataSourceServerDiskSizesKey,
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
		DataSourceServerLocationNameKey,
		DataSourceServerPackageNameKey,
		DataSourceServerTemplateNameKey,
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
