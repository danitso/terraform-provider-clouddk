package main

import (
	"testing"
)

// TestDataSourceDiskInstantiation() tests whether the dataSourceDisk instance can be instantiated.
func TestDataSourceDiskInstantiation(t *testing.T) {
	s := dataSourceDisk()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceDisk")
	}
}

// TestDataSourceDiskSchema() tests the dataSourceDisk schema.
func TestDataSourceDiskSchema(t *testing.T) {
	s := dataSourceDisk()

	idKeys := []string{
		DataSourceDiskDiskIdKey,
		DataSourceDiskIdKey,
	}

	for _, v := range idKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceDisk.Schema: Missing argument \"%s\"", v)
		}

		if s.Schema[v].Required != true {
			t.Fatalf("Error in dataSourceDisk.Schema: Argument \"%s\" is not required", v)
		}
	}

	attributeKeys := []string{
		DataSourceDiskLabelKey,
		DataSourceDiskPrimaryKey,
		DataSourceDiskSizeKey,
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
