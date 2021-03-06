/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package clouddktf

import (
	"testing"
)

// TestDataSourceLocationsInstantiation tests whether the dataSourceLocations instance can be instantiated.
func TestDataSourceLocationsInstantiation(t *testing.T) {
	s := dataSourceLocations()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceLocations")
	}
}

// TestDataSourceLocationsSchema tests the dataSourceLocations schema.
func TestDataSourceLocationsSchema(t *testing.T) {
	s := dataSourceLocations()

	attributeKeys := []string{
		dataSourceLocationsIdsKey,
		dataSourceLocationsNamesKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceLocations.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in dataSourceLocations.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}
