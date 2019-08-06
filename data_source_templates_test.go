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

	attributeKeys := []string{
		dataSourceTemplatesIdsKey,
		dataSourceTemplatesNamesKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceTemplates.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in dataSourceTemplates.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}

// TestDataSourceTemplatesSchemaFilter() tests the dataSourceTemplates.Filter schema.
func TestDataSourceTemplatesSchemaFilter(t *testing.T) {
	s := dataSourceTemplates()

	if s.Schema[dataSourceTemplatesFilterKey] == nil {
		t.Fatalf("Error in dataSourceTemplates.Schema: Missing block \"%s\"", dataSourceTemplatesFilterKey)
	}

	if s.Schema[dataSourceTemplatesFilterKey].Optional != true {
		t.Fatalf("Error in dataSourceTemplates.Schema: Block \"%s\" is not optional", dataSourceTemplatesFilterKey)
	}

	if s.Schema[dataSourceTemplatesFilterKey].Type != schema.TypeList {
		t.Fatalf("Error in dataSourceTemplates.Schema: Block \"%s\" is not a list", dataSourceTemplatesFilterKey)
	}

	if s.Schema[dataSourceTemplatesFilterKey].MaxItems != 1 {
		t.Fatalf("Error in dataSourceTemplates.Schema: Block \"%s\" is not limited to a single definition", dataSourceTemplatesFilterKey)
	}

	if s.Schema[dataSourceTemplatesFilterKey].Elem == nil {
		t.Fatalf("Error in dataSourceTemplates.Schema: Missing element for block \"%s\"", dataSourceTemplatesFilterKey)
	}

	blockElement, blockElementCasted := s.Schema[dataSourceTemplatesFilterKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in dataSourceTemplates.Schema: Element for block \"%s\" is not a pointer to schema.Resource", dataSourceTemplatesFilterKey)
	}

	if blockElement.Schema[dataSourceTemplatesFilterNameKey] == nil {
		t.Fatalf("Error in dataSourceTemplates.Schema.subscriber: Missing argument \"%s\"", dataSourceTemplatesFilterNameKey)
	}

	if blockElement.Schema[dataSourceTemplatesFilterNameKey].Optional != true {
		t.Fatalf("Error in dataSourceTemplates.Schema.subscriber: Argument \"%s\" is not optional", dataSourceTemplatesFilterNameKey)
	}
}
