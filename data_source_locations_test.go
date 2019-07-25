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

	if s.Schema[DataSourceLocationsResult] == nil {
		t.Fatalf("Error in dataSourceLocations.Schema: Missing attribute \"%s\"", DataSourceLocationsResult)
	}

	if s.Schema[DataSourceLocationsResult].Computed != true {
		t.Fatalf("Error in dataSourceLocations.Schema: Attribute \"%s\" is not computed", DataSourceLocationsResult)
	}
}
