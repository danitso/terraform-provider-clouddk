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

	if s.Schema[DataSourcePackagesIdsKey] == nil {
		t.Fatalf("Error in dataSourcePackages.Schema: Missing attribute \"%s\"", DataSourcePackagesIdsKey)
	}

	if s.Schema[DataSourcePackagesIdsKey].Computed != true {
		t.Fatalf("Error in dataSourcePackages.Schema: Attribute \"%s\" is not computed", DataSourcePackagesIdsKey)
	}

	if s.Schema[DataSourcePackagesNamesKey] == nil {
		t.Fatalf("Error in dataSourcePackages.Schema: Missing attribute \"%s\"", DataSourcePackagesNamesKey)
	}

	if s.Schema[DataSourcePackagesNamesKey].Computed != true {
		t.Fatalf("Error in dataSourcePackages.Schema: Attribute \"%s\" is not computed", DataSourcePackagesNamesKey)
	}
}
