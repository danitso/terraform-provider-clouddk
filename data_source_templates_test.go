package main

import (
	"testing"
)

// TestDataSourceTemplatesInstantiation() tests whether the dataSourceTemplates instance can be instantiated.
func TestDataSourceTemplatesInstantiation(t *testing.T) {
	s := dataSourceTemplates()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceTemplates")
	}
}

// TestDataSourceTemplatesSchema() tests the dataSourceTemplates schema.
func TestDataSourceTemplatesSchema(t *testing.T) {
	s := dataSourceTemplates()

	if s.Schema[DataSourceTemplatesResult] == nil {
		t.Fatalf("Error in dataSourceTemplates.Schema: Missing attribute \"%s\"", DataSourceTemplatesResult)
	}

	if s.Schema[DataSourceTemplatesResult].Computed != true {
		t.Fatalf("Error in dataSourceTemplates.Schema: Attribute \"%s\" is not computed", DataSourceTemplatesResult)
	}
}
