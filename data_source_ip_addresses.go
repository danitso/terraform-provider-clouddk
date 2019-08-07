package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	dataSourceIPAddressesAddressesKey           = "addresses"
	dataSourceIPAddressesGatewaysKey            = "gateways"
	dataSourceIPAddressesIDKey                  = "id"
	dataSourceIPAddressesNetmasksKey            = "netmasks"
	dataSourceIPAddressesNetworkInterfaceIdsKey = "network_interface_ids"
	dataSourceIPAddressesNetworksKey            = "networks"
)

// dataSourceIPAddresses retrieves information about IP addresses.
func dataSourceIPAddresses() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourceIPAddressesAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The IP addresses assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceIPAddressesGatewaysKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The gateways assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceIPAddressesIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
			dataSourceIPAddressesNetmasksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The netmasks assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceIPAddressesNetworkInterfaceIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The network interface identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceIPAddressesNetworksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The networks assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},

		Read: dataSourceIPAddressesRead,
	}
}

// dataSourceIPAddressesRead reads information about IP addresses.
func dataSourceIPAddressesRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	id := d.Get(dataSourceIPAddressesIDKey).(string)
	req, err := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/ip-addresses", id), new(bytes.Buffer))

	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the IP addresses - Reason: The API responded with HTTP %s", res.Status)
	}

	ipAddresses := clouddk.IPAddressListBody{}
	err = json.NewDecoder(res.Body).Decode(&ipAddresses)

	if err != nil {
		return err
	}

	addresses := make([]interface{}, len(ipAddresses))
	gateways := make([]interface{}, len(ipAddresses))
	netmasks := make([]interface{}, len(ipAddresses))
	networkInterfaceIds := make([]interface{}, len(ipAddresses))
	networks := make([]interface{}, len(ipAddresses))

	for i, v := range ipAddresses {
		addresses[i] = v.Address
		gateways[i] = v.Gateway
		netmasks[i] = v.Netmask
		networkInterfaceIds[i] = v.NetworkInterfaceIdentifier
		networks[i] = v.Network
	}

	d.SetId(id)

	d.Set(dataSourceIPAddressesAddressesKey, addresses)
	d.Set(dataSourceIPAddressesGatewaysKey, gateways)
	d.Set(dataSourceIPAddressesNetmasksKey, netmasks)
	d.Set(dataSourceIPAddressesNetworkInterfaceIdsKey, networkInterfaceIds)
	d.Set(dataSourceIPAddressesNetworksKey, networks)

	return nil
}
