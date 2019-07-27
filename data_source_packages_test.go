package main

import (
	"testing"
)

// TestDataSourcePackagesInstantiation() tests whether the dataSourcePackages instance can be instantiated.
func TestDataSourcePackagesInstantiation(t *testing.T) {
	s := dataSourcePackages()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourcePackages")
	}
}

// TestDataSourcePackagesSchema() tests the dataSourcePackages schema.
func TestDataSourcePackagesSchema(t *testing.T) {
	s := dataSourcePackages()

	attributeKeys := []string{
		DataSourcePackagesIdsKey,
		DataSourcePackagesNamesKey,
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
