package main

import (
	"testing"
)

// TestDataSourceLocationsInstantiation() tests whether the dataSourceLocations instance can be instantiated.
func TestDataSourceLocationsInstantiation(t *testing.T) {
	s := dataSourceLocations()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceLocations")
	}
}

// TestDataSourceLocationsSchema() tests the dataSourceLocations schema.
func TestDataSourceLocationsSchema(t *testing.T) {
	s := dataSourceLocations()

	if s.Schema[DataSourceLocationsIdsKey] == nil {
		t.Fatalf("Error in dataSourceLocations.Schema: Missing attribute \"%s\"", DataSourceLocationsIdsKey)
	}

	if s.Schema[DataSourceLocationsIdsKey].Computed != true {
		t.Fatalf("Error in dataSourceLocations.Schema: Attribute \"%s\" is not computed", DataSourceLocationsIdsKey)
	}

	if s.Schema[DataSourceLocationsNamesKey] == nil {
		t.Fatalf("Error in dataSourceLocations.Schema: Missing attribute \"%s\"", DataSourceLocationsNamesKey)
	}

	if s.Schema[DataSourceLocationsNamesKey].Computed != true {
		t.Fatalf("Error in dataSourceLocations.Schema: Attribute \"%s\" is not computed", DataSourceLocationsNamesKey)
	}
}
