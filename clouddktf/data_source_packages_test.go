/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package clouddktf

import (
	"testing"
)

// TestDataSourcePackagesInstantiation tests whether the dataSourcePackages instance can be instantiated.
func TestDataSourcePackagesInstantiation(t *testing.T) {
	s := dataSourcePackages()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourcePackages")
	}
}

// TestDataSourcePackagesSchema tests the dataSourcePackages schema.
func TestDataSourcePackagesSchema(t *testing.T) {
	s := dataSourcePackages()

	attributeKeys := []string{
		dataSourcePackagesIdsKey,
		dataSourcePackagesNamesKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourcePackages.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in dataSourcePackages.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}
