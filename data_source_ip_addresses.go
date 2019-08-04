package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceIPAddressesAddressesKey = "addresses"
const DataSourceIPAddressesGatewaysKey = "gateways"
const DataSourceIPAddressesIdKey = "id"
const DataSourceIPAddressesNetmasksKey = "netmasks"
const DataSourceIPAddressesNetworkInterfaceIdsKey = "network_interface_ids"
const DataSourceIPAddressesNetworksKey = "networks"

// dataSourceIPAddresses() retrieves information about IP addresses.
func dataSourceIPAddresses() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceIPAddressesAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The IP addresses assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceIPAddressesGatewaysKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The gateways assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceIPAddressesIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
			DataSourceIPAddressesNetmasksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The netmasks assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceIPAddressesNetworkInterfaceIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The network interface identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceIPAddressesNetworksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The networks assigned to the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},

		Read: dataSourceIPAddressesRead,
	}
}

// dataSourceIPAddressesRead() reads information about IP addresses.
func dataSourceIPAddressesRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	id := d.Get(DataSourceIPAddressesIdKey).(string)
	req, reqErr := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/ip-addresses", id), new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the IP addresses - Reason: The API responded with HTTP %s", res.Status)
	}

	ipAddresses := clouddk.IPAddressListBody{}
	json.NewDecoder(res.Body).Decode(&ipAddresses)

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

	d.Set(DataSourceIPAddressesAddressesKey, addresses)
	d.Set(DataSourceIPAddressesGatewaysKey, gateways)
	d.Set(DataSourceIPAddressesNetmasksKey, netmasks)
	d.Set(DataSourceIPAddressesNetworkInterfaceIdsKey, networkInterfaceIds)
	d.Set(DataSourceIPAddressesNetworksKey, networks)

	return nil
}
