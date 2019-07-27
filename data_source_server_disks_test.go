package main

import (
	"testing"
)

// TestDataSourceServerDisksInstantiation() tests whether the dataSourceServer instance can be instantiated.
func TestDataSourceServerDisksInstantiation(t *testing.T) {
	s := dataSourceServerDisks()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceServerDisks")
	}
}

// TestDataSourceServerDisksSchema() tests the dataSourceServer schema.
func TestDataSourceServerDisksSchema(t *testing.T) {
	s := dataSourceServerDisks()

	if s.Schema[DataSourceServerDisksIdKey] == nil {
		t.Fatalf("Error in dataSourceServerDisks.Schema: Missing argument \"%s\"", DataSourceServerIdKey)
	}

	if s.Schema[DataSourceServerDisksIdKey].Required != true {
		t.Fatalf("Error in dataSourceServerDisks.Schema: Argument \"%s\" is not required", DataSourceServerIdKey)
	}

	attributeKeys := []string{
		DataSourceServerDisksIdsKey,
		DataSourceServerDisksLabelsKey,
		DataSourceServerDisksPrimaryKey,
		DataSourceServerDisksSizesKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceServerDisks.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in dataSourceServerDisks.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}
