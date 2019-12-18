/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package clouddktf

import (
	"testing"
)

// TestDataSourceIPAddressesInstantiation tests whether the dataSourceIPAddresses instance can be instantiated.
func TestDataSourceIPAddressesInstantiation(t *testing.T) {
	s := dataSourceIPAddresses()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceIPAddresses")
	}
}

// TestDataSourceIPAddressesSchema tests the dataSourceIPAddresses schema.
func TestDataSourceIPAddressesSchema(t *testing.T) {
	s := dataSourceIPAddresses()

	if s.Schema[dataSourceIPAddressesIDKey] == nil {
		t.Fatalf("Error in dataSourceIPAddresses.Schema: Missing argument \"%s\"", dataSourceIPAddressesIDKey)
	}

	if s.Schema[dataSourceIPAddressesIDKey].Required != true {
		t.Fatalf("Error in dataSourceIPAddresses.Schema: Argument \"%s\" is not required", dataSourceIPAddressesIDKey)
	}

	attributeKeys := []string{
		dataSourceIPAddressesAddressesKey,
		dataSourceIPAddressesGatewaysKey,
		dataSourceIPAddressesNetmasksKey,
		dataSourceIPAddressesNetworkInterfaceIdsKey,
		dataSourceIPAddressesNetworksKey,
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
