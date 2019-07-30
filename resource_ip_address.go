package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

const ResourceIPAddressAddressKey = "address"
const ResourceIPAddressGatewayKey = "gateway"
const ResourceIPAddressNetmaskKey = "netmask"
const ResourceIPAddressNetworkKey = "network"
const ResourceIPAddressNetworkInterfaceIdKey = "network_interface_id"
const ResourceIPAddressServerIdKey = "server_id"

// resourceIPAddress() manages an IP address.
func resourceIPAddress() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			ResourceIPAddressAddressKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP address",
			},
			ResourceIPAddressGatewayKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The gateway address",
			},
			ResourceIPAddressNetmaskKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The netmask",
			},
			ResourceIPAddressNetworkKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The network address",
			},
			ResourceIPAddressNetworkInterfaceIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The network interface id",
			},
			ResourceIPAddressServerIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
		},

		Create: resourceIPAddressCreate,
		Read:   resourceIPAddressRead,
		Delete: resourceIPAddressDelete,
	}
}

// resourceIPAddressCreate() creates an IP address.
func resourceIPAddressCreate(d *schema.ResourceData, m interface{}) error {
	serverId := d.Get(ResourceIPAddressServerIdKey).(string)

	// We need to wait for transactions to end before proceeding.
	transactionsErr := resourceServerWaitForTransactions(d, m, serverId, 60, 10)

	if transactionsErr != nil {
		return transactionsErr
	}

	// We should now be able to create the IP address without any issues.
	clientSettings := m.(ClientSettings)

	res, resErr := doClientRequest(&clientSettings, "POST", fmt.Sprintf("cloudservers/%s/ip-addresses", serverId), new(bytes.Buffer), []int{200}, 60, 10)

	if resErr != nil {
		return resErr
	}

	ipAddresses := IPAddressListBody{}
	json.NewDecoder(res.Body).Decode(&ipAddresses)

	d.SetId(ipAddresses[len(ipAddresses)-1].Address)

	d.Set(ResourceIPAddressAddressKey, ipAddresses[len(ipAddresses)-1].Address)
	d.Set(ResourceIPAddressGatewayKey, ipAddresses[len(ipAddresses)-1].Gateway)
	d.Set(ResourceIPAddressNetmaskKey, ipAddresses[len(ipAddresses)-1].Netmask)
	d.Set(ResourceIPAddressNetworkKey, ipAddresses[len(ipAddresses)-1].Network)
	d.Set(ResourceIPAddressNetworkInterfaceIdKey, ipAddresses[len(ipAddresses)-1].NetworkInterfaceIdentifier)

	return nil
}

// resourceIPAddressRead reads information about an existing IP address.
func resourceIPAddressRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	address := d.Id()
	serverId := d.Get(ResourceIPAddressServerIdKey).(string)

	req, reqErr := getClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/ip-addresses", serverId), new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	} else if res.StatusCode != 200 {
		if res.StatusCode == 404 {
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Failed to read the IP address information - Reason: The API responded with HTTP %s", res.Status)
	}

	ipAddresses := IPAddressListBody{}
	json.NewDecoder(res.Body).Decode(&ipAddresses)

	for _, v := range ipAddresses {
		if v.Address == address {
			d.Set(ResourceIPAddressAddressKey, v.Address)
			d.Set(ResourceIPAddressGatewayKey, v.Gateway)
			d.Set(ResourceIPAddressNetmaskKey, v.Netmask)
			d.Set(ResourceIPAddressNetworkKey, v.Network)
			d.Set(ResourceIPAddressNetworkInterfaceIdKey, v.NetworkInterfaceIdentifier)

			return nil
		}
	}

	d.SetId("")

	return nil
}

// resourceIPAddressDelete deletes an existing IP address.
func resourceIPAddressDelete(d *schema.ResourceData, m interface{}) error {
	serverId := d.Get(ResourceIPAddressServerIdKey).(string)

	// We need to wait for transactions to end before proceeding.
	transactionsErr := resourceServerWaitForTransactions(d, m, serverId, 60, 10)

	if transactionsErr != nil {
		return transactionsErr
	}

	// We should now be able to delete the IP address without any issues.
	clientSettings := m.(ClientSettings)

	address := d.Id()

	_, err := doClientRequest(&clientSettings, "DELETE", fmt.Sprintf("cloudservers/%s/ip-addresses?address=%s", serverId, address), new(bytes.Buffer), []int{200, 404}, 60, 10)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
