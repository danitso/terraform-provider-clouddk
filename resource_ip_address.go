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
	resourceIPAddressAddressKey            = "address"
	resourceIPAddressGatewayKey            = "gateway"
	resourceIPAddressNetmaskKey            = "netmask"
	resourceIPAddressNetworkKey            = "network"
	resourceIPAddressNetworkInterfaceIDKey = "network_interface_id"
	resourceIPAddressServerIDKey           = "server_id"
)

// resourceIPAddress manages an IP address.
func resourceIPAddress() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			resourceIPAddressAddressKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP address",
			},
			resourceIPAddressGatewayKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The gateway address",
			},
			resourceIPAddressNetmaskKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The netmask",
			},
			resourceIPAddressNetworkKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The network address",
			},
			resourceIPAddressNetworkInterfaceIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The network interface id",
			},
			resourceIPAddressServerIDKey: &schema.Schema{
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

// resourceIPAddressCreate creates an IP address.
func resourceIPAddressCreate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	serverID := d.Get(resourceIPAddressServerIDKey).(string)

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	err := resourceServerLock(d, m, serverID)

	if err != nil {
		return err
	}

	res, err := clouddk.DoClientRequest(&clientSettings, "POST", fmt.Sprintf("cloudservers/%s/ip-addresses", serverID), new(bytes.Buffer), []int{200}, 60, 10)

	if err != nil {
		resourceServerUnlock(d, m, serverID)

		return err
	}

	err = resourceServerUnlock(d, m, serverID)

	if err != nil {
		return err
	}

	ipAddresses := clouddk.IPAddressListBody{}
	err = json.NewDecoder(res.Body).Decode(&ipAddresses)

	if err != nil {
		return err
	}

	d.SetId(ipAddresses[len(ipAddresses)-1].Address)

	d.Set(resourceIPAddressAddressKey, ipAddresses[len(ipAddresses)-1].Address)
	d.Set(resourceIPAddressGatewayKey, ipAddresses[len(ipAddresses)-1].Gateway)
	d.Set(resourceIPAddressNetmaskKey, ipAddresses[len(ipAddresses)-1].Netmask)
	d.Set(resourceIPAddressNetworkKey, ipAddresses[len(ipAddresses)-1].Network)
	d.Set(resourceIPAddressNetworkInterfaceIDKey, ipAddresses[len(ipAddresses)-1].NetworkInterfaceIdentifier)

	return nil
}

// resourceIPAddressRead reads information about an existing IP address.
func resourceIPAddressRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	address := d.Id()
	serverID := d.Get(resourceIPAddressServerIDKey).(string)

	req, err := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/ip-addresses", serverID), new(bytes.Buffer))

	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		if res.StatusCode == 404 {
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Failed to read the IP address information - Reason: The API responded with HTTP %s", res.Status)
	}

	ipAddresses := clouddk.IPAddressListBody{}
	err = json.NewDecoder(res.Body).Decode(&ipAddresses)

	if err != nil {
		return err
	}

	for _, v := range ipAddresses {
		if v.Address == address {
			d.Set(resourceIPAddressAddressKey, v.Address)
			d.Set(resourceIPAddressGatewayKey, v.Gateway)
			d.Set(resourceIPAddressNetmaskKey, v.Netmask)
			d.Set(resourceIPAddressNetworkKey, v.Network)
			d.Set(resourceIPAddressNetworkInterfaceIDKey, v.NetworkInterfaceIdentifier)

			return nil
		}
	}

	d.SetId("")

	return nil
}

// resourceIPAddressDelete deletes an existing IP address.
func resourceIPAddressDelete(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	serverID := d.Get(resourceIPAddressServerIDKey).(string)
	address := d.Id()

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	err := resourceServerLock(d, m, serverID)

	if err != nil {
		return err
	}

	_, err = clouddk.DoClientRequest(&clientSettings, "DELETE", fmt.Sprintf("cloudservers/%s/ip-addresses?address=%s", serverID, address), new(bytes.Buffer), []int{200, 404}, 60, 10)

	if err != nil {
		resourceServerUnlock(d, m, serverID)

		return err
	}

	err = resourceServerUnlock(d, m, serverID)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
