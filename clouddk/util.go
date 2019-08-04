package clouddk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"
)

// GetClientRequestObject() returns a new HTTP request object.
func GetClientRequestObject(settings *ClientSettings, method string, path string, body io.Reader) (*http.Request, error) {
	req, reqErr := http.NewRequest(method, fmt.Sprintf("%s/%s", settings.Endpoint, path), body)

	if reqErr != nil {
		return nil, reqErr
	}

	req.Header.Set("User-Agent", fmt.Sprintf("Go/%s", runtime.Version()))
	req.Header.Set("X-Api-Key", settings.Key)

	return req, nil
}

// DoClientRequest() performs a HTTP request and does so multiple times, if required.
func DoClientRequest(settings *ClientSettings, method string, path string, body *bytes.Buffer, successCodes []int, retryLimit int, retryDelay int) (*http.Response, error) {
	timeDelay := int64(retryDelay)
	timeMax := float64(retryLimit * retryDelay)
	timeStart := time.Now()
	timeElapsed := timeStart.Sub(timeStart)

	var response *http.Response
	var responseError error

	bodyString := body.String()
	errorMessage := ""

	for timeElapsed.Seconds() < timeMax {
		if int64(timeElapsed.Seconds())%timeDelay == 0 {
			log.Printf("[DEBUG] Querying the API - Method: %s - Path: %s", method, path)

			requestBody := bytes.NewBufferString(bodyString)
			request, requestError := GetClientRequestObject(settings, method, path, requestBody)

			if requestError != nil {
				return nil, requestError
			}

			if requestBody.Len() > 0 {
				request.Header.Set("Content-Type", "application/json")
				log.Printf("[DEBUG] Adding body to request - Method: %s - Path: %s - Content-Type: %s - Content-Length: %d - Body: %s", method, path, request.Header.Get("Content-Type"), requestBody.Len(), bodyString)
			} else if method == "POST" || method == "PUT" {
				log.Printf("[DEBUG] WARNING: No request body specified - Method: %s - Path: %s", method, path)
			}

			client := &http.Client{}
			response, responseError = client.Do(request)

			if responseError != nil {
				return response, responseError
			}

			for _, v := range successCodes {
				if response.StatusCode == v {
					log.Printf("[DEBUG] The API query was successful - Method: %s - Path: %s", method, path)

					return response, nil
				}
			}

			errorBody := ErrorBody{}
			json.NewDecoder(response.Body).Decode(&errorBody)

			if len(errorBody.Message) > 0 {
				errorMessage = fmt.Sprintf("%s (HTTP %d)", errorBody.Message, response.StatusCode)
			} else {
				errorMessage = fmt.Sprintf("HTTP %s", response.Status)
			}

			if response.StatusCode != 500 {
				if response.StatusCode != 400 || !strings.Contains(errorBody.Message, "CloudServer that is not yet built") {
					break
				}
			}

			log.Printf("[DEBUG] Failed to query the API - Reason: %s - Method: %s - Path: %s", errorMessage, method, path)
			time.Sleep(1 * time.Second)
		}

		time.Sleep(200 * time.Millisecond)

		timeElapsed = time.Now().Sub(timeStart)
	}

	return response, fmt.Errorf("Failed to query the API - Reason: %s - Method: %s - Path: %s", errorMessage, method, path)
}
