package main

import (
	"testing"
)

// TestDataSourceServerPackagesInstantiation() tests whether the dataSourceServerPackages instance can be instantiated.
func TestDataSourceServerPackagesInstantiation(t *testing.T) {
	s := dataSourceServerPackages()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceServerPackages")
	}
}

// TestDataSourceServerPackagesSchema() tests the dataSourceServerPackages schema.
func TestDataSourceServerPackagesSchema(t *testing.T) {
	s := dataSourceServerPackages()

	if s.Schema[DataSourceServerPackagesIdsKey] == nil {
		t.Fatalf("Error in dataSourceServerPackages.Schema: Missing attribute \"%s\"", DataSourceServerPackagesIdsKey)
	}

	if s.Schema[DataSourceServerPackagesIdsKey].Computed != true {
		t.Fatalf("Error in dataSourceServerPackages.Schema: Attribute \"%s\" is not computed", DataSourceServerPackagesIdsKey)
	}

	if s.Schema[DataSourceServerPackagesNamesKey] == nil {
		t.Fatalf("Error in dataSourceServerPackages.Schema: Missing attribute \"%s\"", DataSourceServerPackagesNamesKey)
	}

	if s.Schema[DataSourceServerPackagesNamesKey].Computed != true {
		t.Fatalf("Error in dataSourceServerPackages.Schema: Attribute \"%s\" is not computed", DataSourceServerPackagesNamesKey)
	}
}
