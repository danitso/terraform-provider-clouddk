package main

import (
	"fmt"
	"io"
	"net/http"
)

// getClientRequestObject() returns a new HTTP request object.
func getClientRequestObject(settings *ClientSettings, method string, path string, body io.Reader) (*http.Request, error) {
	req, reqErr := http.NewRequest(method, fmt.Sprintf("%s/%s", settings.Endpoint, path), body)

	if reqErr != nil {
		return nil, reqErr
	}

	req.Header.Set("X-Api-Key", settings.Key)

	return req, nil
}
