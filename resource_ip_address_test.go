package main

import (
	"testing"
)

// TestResourceIPAddressInstantiation() tests whether the resourceIPAddress instance can be instantiated.
func TestResourceIPAddressInstantiation(t *testing.T) {
	s := resourceIPAddress()

	if s == nil {
		t.Fatalf("Cannot instantiate resourceIPAddress")
	}
}

// TestResourceIPAddressSchema() tests the resourceIPAddress schema.
func TestResourceIPAddressSchema(t *testing.T) {
	s := resourceIPAddress()

	requiredKeys := []string{
		ResourceIPAddressServerIdKey,
	}

	for _, v := range requiredKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in resourceIPAddress.Schema: Missing argument \"%s\"", v)
		}

		if s.Schema[v].Required != true {
			t.Fatalf("Error in resourceIPAddress.Schema: Argument \"%s\" is not required", v)
		}
	}

	attributeKeys := []string{
		ResourceIPAddressAddressKey,
		ResourceIPAddressGatewayKey,
		ResourceIPAddressNetmaskKey,
		ResourceIPAddressNetworkKey,
		ResourceIPAddressNetworkInterfaceIdKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in resourceIPAddress.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in resourceIPAddress.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}
