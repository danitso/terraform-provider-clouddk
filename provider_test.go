package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
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

	if s.Schema[ProviderConfigurationEndpoint] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\"", ProviderConfigurationEndpoint)
	}

	if s.Schema[ProviderConfigurationEndpoint].Optional != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" is not optional", ProviderConfigurationEndpoint)
	}

	if s.Schema[ProviderConfigurationEndpoint].Type != schema.TypeString {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" is not a string", ProviderConfigurationEndpoint)
	}

	if s.Schema[ProviderConfigurationKey] == nil {
		t.Fatalf("Error in Provider.Schema: Missing argument \"%s\"", ProviderConfigurationKey)
	}

	if s.Schema[ProviderConfigurationKey].Required != true {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" is not required", ProviderConfigurationKey)
	}

	if s.Schema[ProviderConfigurationKey].Type != schema.TypeString {
		t.Fatalf("Error in Provider.Schema: Argument \"%s\" is not a string", ProviderConfigurationKey)
	}
}
