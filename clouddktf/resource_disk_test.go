/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package clouddktf

import (
	"testing"
)

// TestResourceDiskInstantiation tests whether the resourceDisk instance can be instantiated.
func TestResourceDiskInstantiation(t *testing.T) {
	s := resourceDisk()

	if s == nil {
		t.Fatalf("Cannot instantiate resourceDisk")
	}
}

// TestResourceDiskSchema tests the resourceDisk schema.
func TestResourceDiskSchema(t *testing.T) {
	s := resourceDisk()

	requiredKeys := []string{
		dataSourceDiskLabelKey,
		dataSourceDiskServerIDKey,
		dataSourceDiskSizeKey,
	}

	for _, v := range requiredKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in resourceDisk.Schema: Missing argument \"%s\"", v)
		}

		if s.Schema[v].Required != true {
			t.Fatalf("Error in resourceDisk.Schema: Argument \"%s\" is not required", v)
		}
	}

	attributeKeys := []string{
		dataSourceDiskPrimaryKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceDisk.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in dataSourceDisk.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}
