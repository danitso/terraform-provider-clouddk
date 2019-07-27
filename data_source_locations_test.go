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

	attributeKeys := []string{
		DataSourceLocationsIdsKey,
		DataSourceLocationsNamesKey,
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
