package main

// ClientSettings describes the client settings.
type ClientSettings struct {
	Endpoint string
	Key      string
}

// LocationBody describes a datacenter location payload.
type LocationBody struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

// LocationListBody describes a datacenter location list payload.
type LocationListBody []LocationBody
