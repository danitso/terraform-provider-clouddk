package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
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

	if s.Schema[DataSourceTemplatesIdsKey] == nil {
		t.Fatalf("Error in dataSourceLocations.Schema: Missing attribute \"%s\"", DataSourceTemplatesIdsKey)
	}

	if s.Schema[DataSourceTemplatesIdsKey].Computed != true {
		t.Fatalf("Error in dataSourceLocations.Schema: Attribute \"%s\" is not computed", DataSourceTemplatesIdsKey)
	}

	if s.Schema[DataSourceTemplatesNamesKey] == nil {
		t.Fatalf("Error in dataSourceLocations.Schema: Missing attribute \"%s\"", DataSourceTemplatesNamesKey)
	}

	if s.Schema[DataSourceTemplatesNamesKey].Computed != true {
		t.Fatalf("Error in dataSourceLocations.Schema: Attribute \"%s\" is not computed", DataSourceTemplatesNamesKey)
	}
}

// TestDataSourceTemplatesSchemaFilter() tests the dataSourceTemplates.Filter schema.
func TestDataSourceTemplatesSchemaFilter(t *testing.T) {
	s := dataSourceTemplates()

	if s.Schema[DataSourceTemplatesFilterKey] == nil {
		t.Fatalf("Error in dataSourceTemplates.Schema: Missing block \"%s\"", DataSourceTemplatesFilterKey)
	}

	if s.Schema[DataSourceTemplatesFilterKey].Optional != true {
		t.Fatalf("Error in dataSourceTemplates.Schema: Block \"%s\" is not optional", DataSourceTemplatesFilterKey)
	}

	if s.Schema[DataSourceTemplatesFilterKey].Type != schema.TypeList {
		t.Fatalf("Error in dataSourceTemplates.Schema: Block \"%s\" is not a list", DataSourceTemplatesFilterKey)
	}

	if s.Schema[DataSourceTemplatesFilterKey].MaxItems != 1 {
		t.Fatalf("Error in dataSourceTemplates.Schema: Block \"%s\" is not limited to a single definition", DataSourceTemplatesFilterKey)
	}

	if s.Schema[DataSourceTemplatesFilterKey].Elem == nil {
		t.Fatalf("Error in dataSourceTemplates.Schema: Missing element for block \"%s\"", DataSourceTemplatesFilterKey)
	}

	blockElement, blockElementCasted := s.Schema[DataSourceTemplatesFilterKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in dataSourceTemplates.Schema: Element for block \"%s\" is not a pointer to schema.Resource", DataSourceTemplatesFilterKey)
	}

	if blockElement.Schema[DataSourceTemplatesFilterNameKey] == nil {
		t.Fatalf("Error in dataSourceTemplates.Schema.subscriber: Missing argument \"%s\"", DataSourceTemplatesFilterNameKey)
	}

	if blockElement.Schema[DataSourceTemplatesFilterNameKey].Optional != true {
		t.Fatalf("Error in dataSourceTemplates.Schema.subscriber: Argument \"%s\" is not optional", DataSourceTemplatesFilterNameKey)
	}
}
