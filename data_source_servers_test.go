/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

// TestDataSourceServersInstantiation tests whether the dataSourceServers instance can be instantiated.
func TestDataSourceServersInstantiation(t *testing.T) {
	s := dataSourceServers()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceServers")
	}
}

// TestDataSourceServersSchema tests the dataSourceServers schema.
func TestDataSourceServersSchema(t *testing.T) {
	s := dataSourceServers()

	attributeKeys := []string{
		dataSourceServersHostnamesKey,
		dataSourceServersIdsKey,
		dataSourceServersLabelsKey,
		dataSourceServersLocationIdsKey,
		dataSourceServersLocationNamesKey,
		dataSourceServersPackageIdsKey,
		dataSourceServersPackageNamesKey,
		dataSourceServersTemplateIdsKey,
		dataSourceServersTemplateNamesKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceServers.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in dataSourceServers.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}

// TestDataSourceServersSchemaFilter tests the dataSourceServers.Filter schema.
func TestDataSourceServersSchemaFilter(t *testing.T) {
	s := dataSourceServers()

	if s.Schema[dataSourceServersFilterKey] == nil {
		t.Fatalf("Error in dataSourceServers.Schema: Missing block \"%s\"", dataSourceServersFilterKey)
	}

	if s.Schema[dataSourceServersFilterKey].Optional != true {
		t.Fatalf("Error in dataSourceServers.Schema: Block \"%s\" is not optional", dataSourceServersFilterKey)
	}

	if s.Schema[dataSourceServersFilterKey].Type != schema.TypeList {
		t.Fatalf("Error in dataSourceServers.Schema: Block \"%s\" is not a list", dataSourceServersFilterKey)
	}

	if s.Schema[dataSourceServersFilterKey].MaxItems != 1 {
		t.Fatalf("Error in dataSourceServers.Schema: Block \"%s\" is not limited to a single definition", dataSourceServersFilterKey)
	}

	if s.Schema[dataSourceServersFilterKey].Elem == nil {
		t.Fatalf("Error in dataSourceServers.Schema: Missing element for block \"%s\"", dataSourceServersFilterKey)
	}

	blockElement, blockElementCasted := s.Schema[dataSourceServersFilterKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in dataSourceServers.Schema: Element for block \"%s\" is not a pointer to schema.Resource", dataSourceServersFilterKey)
	}

	if blockElement.Schema[dataSourceServersFilterHostnameKey] == nil {
		t.Fatalf("Error in dataSourceServers.Schema.subscriber: Missing argument \"%s\"", dataSourceServersFilterHostnameKey)
	}

	if blockElement.Schema[dataSourceServersFilterHostnameKey].Optional != true {
		t.Fatalf("Error in dataSourceServers.Schema.subscriber: Argument \"%s\" is not optional", dataSourceServersFilterHostnameKey)
	}
}
