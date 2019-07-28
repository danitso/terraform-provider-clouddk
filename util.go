package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"time"

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

// doClientRequest() performs a HTTP request and does so multiple times, if required.
func doClientRequest(r *http.Request, successCodes []int, retryLimit int, retryDelay int) (*http.Response, error) {
	timeDelay := int64(retryDelay)
	timeMax := float64(retryLimit * retryDelay)
	timeStart := time.Now()
	timeElapsed := timeStart.Sub(timeStart)

	var response *http.Response
	var responseError error

	responseStatus := ""

	for timeElapsed.Seconds() < timeMax {
		if int64(timeElapsed.Seconds())%timeDelay == 0 {
			log.Printf("[DEBUG] Querying the API - URI: %s", r.RequestURI)

			client := &http.Client{}
			response, responseError = client.Do(r)

			if responseError != nil {
				return response, responseError
			}

			for _, v := range successCodes {
				if response.StatusCode == v {
					log.Printf("[DEBUG] The API query was successful - URI: %s", r.RequestURI)

					return response, nil
				}
			}

			responseStatus = response.Status

			if response.StatusCode != 500 {
				break
			}

			log.Printf("[DEBUG] Failed to query the API - Reason: HTTP %s - URI: %s", responseStatus, r.RequestURI)
			time.Sleep(1 * time.Second)
		}

		time.Sleep(200 * time.Millisecond)

		timeElapsed = time.Now().Sub(timeStart)
	}

	return response, fmt.Errorf("Failed to query the API - Reason: HTTP %s - URI: %s", responseStatus, r.RequestURI)
}
