package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/hashicorp/terraform/version"
)

// getClientRequestObject() returns a new HTTP request object.
func getClientRequestObject(settings *ClientSettings, method string, path string, body io.Reader) (*http.Request, error) {
	req, reqErr := http.NewRequest(method, fmt.Sprintf("%s/%s", settings.Endpoint, path), body)

	if reqErr != nil {
		return nil, reqErr
	}

	req.Header.Set("User-Agent", fmt.Sprintf(
		"%s/%s Go/%s Terraform-Library/%s",
		TerraformProviderName,
		TerraformProviderVersion,
		runtime.Version(),
		version.Version,
	))
	req.Header.Set("X-Api-Key", settings.Key)

	return req, nil
}
