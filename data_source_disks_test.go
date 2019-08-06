package main

import (
	"testing"
)

// TestDataSourceDisksInstantiation tests whether the dataSourceDisks instance can be instantiated.
func TestDataSourceDisksInstantiation(t *testing.T) {
	s := dataSourceDisks()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceDisks")
	}
}

// TestDataSourceDisksSchema tests the dataSourceDisks schema.
func TestDataSourceDisksSchema(t *testing.T) {
	s := dataSourceDisks()

	if s.Schema[dataSourceDisksIDKey] == nil {
		t.Fatalf("Error in dataSourceDisks.Schema: Missing argument \"%s\"", dataSourceDisksIDKey)
	}

	if s.Schema[dataSourceDisksIDKey].Required != true {
		t.Fatalf("Error in dataSourceDisks.Schema: Argument \"%s\" is not required", dataSourceDisksIDKey)
	}

	attributeKeys := []string{
		dataSourceDisksIdsKey,
		dataSourceDisksLabelsKey,
		dataSourceDisksPrimaryKey,
		dataSourceDisksSizesKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceDisks.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in dataSourceDisks.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}
