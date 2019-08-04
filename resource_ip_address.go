package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
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
	clientSettings := m.(clouddk.ClientSettings)

	serverId := d.Get(ResourceIPAddressServerIdKey).(string)

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	lockErr := resourceServerLock(d, m, serverId)

	if lockErr != nil {
		return lockErr
	}

	res, resErr := clouddk.DoClientRequest(&clientSettings, "POST", fmt.Sprintf("cloudservers/%s/ip-addresses", serverId), new(bytes.Buffer), []int{200}, 60, 10)

	if resErr != nil {
		resourceServerUnlock(d, m, serverId)

		return resErr
	}

	lockErr = resourceServerUnlock(d, m, serverId)

	if lockErr != nil {
		return lockErr
	}

	ipAddresses := clouddk.IPAddressListBody{}
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
	clientSettings := m.(clouddk.ClientSettings)

	address := d.Id()
	serverId := d.Get(ResourceIPAddressServerIdKey).(string)

	req, reqErr := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/ip-addresses", serverId), new(bytes.Buffer))

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

	ipAddresses := clouddk.IPAddressListBody{}
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
	clientSettings := m.(clouddk.ClientSettings)

	serverId := d.Get(ResourceIPAddressServerIdKey).(string)
	address := d.Id()

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	lockErr := resourceServerLock(d, m, serverId)

	if lockErr != nil {
		return lockErr
	}

	_, err := clouddk.DoClientRequest(&clientSettings, "DELETE", fmt.Sprintf("cloudservers/%s/ip-addresses?address=%s", serverId, address), new(bytes.Buffer), []int{200, 404}, 60, 10)

	if err != nil {
		resourceServerUnlock(d, m, serverId)

		return err
	}

	lockErr = resourceServerUnlock(d, m, serverId)

	if lockErr != nil {
		return lockErr
	}

	d.SetId("")

	return nil
}
