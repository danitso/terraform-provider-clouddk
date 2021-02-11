/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package clouddktf

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// TestProviderInstantiation() tests whether the Provider instance can be instantiated.
func TestProviderInstantiation(t *testing.T) {
	s := Provider()

	if s == nil {
		t.Fatalf("Cannot instantiate Provider")
	}
}

// TestProviderSchema() tests the Provider schema.
func TestProviderSchema(t *testing.T) {
	s := Provider()

	if s.Schema[providerConfigurationEndpoint] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\"", providerConfigurationEndpoint)
	}

	if s.Schema[providerConfigurationEndpoint].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" is not optional", providerConfigurationEndpoint)
	}

	if s.Schema[providerConfigurationEndpoint].Type != schema.TypeString {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" is not a string", providerConfigurationEndpoint)
	}

	if s.Schema[providerConfigurationKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\"", providerConfigurationKey)
	}

	if s.Schema[providerConfigurationKey].Required != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" is not required", providerConfigurationKey)
	}

	if s.Schema[providerConfigurationKey].Type != schema.TypeString {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" is not a string", providerConfigurationKey)
	}
}
