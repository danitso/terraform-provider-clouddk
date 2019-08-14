/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform/helper/schema"
)

// resourceFirewallRule manages a firewall rule.
func resourceFirewallRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourceFirewallRuleAddressKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The CIDR block for the firewall rule",
			},
			dataSourceFirewallRuleCommandKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The command for the firewall rule",
			},
			dataSourceFirewallRuleNetworkInterfaceIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The network interface identifier",
				ForceNew:    true,
			},
			dataSourceFirewallRulePortKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The port for the firewall rule",
			},
			dataSourceFirewallRuleProtocolKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The protocol for the firewall rule",
			},
			dataSourceFirewallRuleServerIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server identifier",
				ForceNew:    true,
			},
		},

		Create: resourceFirewallRuleCreate,
		Read:   resourceFirewallRuleRead,
		Update: resourceFirewallRuleUpdate,
		Delete: resourceFirewallRuleDelete,
	}
}

// resourceFirewallRuleCreate creates a firewall rule.
func resourceFirewallRuleCreate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	serverID := d.Get(dataSourceFirewallRuleServerIDKey).(string)
	networkInterfaceID := d.Get(dataSourceFirewallRuleNetworkInterfaceIDKey).(string)
	address := strings.Split(d.Get(dataSourceFirewallRuleAddressKey).(string), "/")

	if len(address) != 2 {
		return fmt.Errorf("Invalid address '%s' for firewall rule (must be defined as x.x.x.x/x)", d.Get(dataSourceFirewallRuleAddressKey).(string))
	}

	bits, err := strconv.Atoi(address[1])

	if err != nil {
		return fmt.Errorf("Invalid address '%s' for firewall rule (%s)", d.Get(dataSourceFirewallRuleAddressKey).(string), err.Error())
	}

	body := clouddk.FirewallRuleCreateBody{
		Command:  d.Get(dataSourceFirewallRuleCommandKey).(string),
		Protocol: d.Get(dataSourceFirewallRuleProtocolKey).(string),
		Address:  address[0],
		Bits:     clouddk.CustomInt(bits),
		Port:     d.Get(dataSourceFirewallRulePortKey).(string),
	}

	reqBody := new(bytes.Buffer)
	err = json.NewEncoder(reqBody).Encode(body)

	if err != nil {
		return err
	}

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	err = resourceServerLock(d, m, serverID)

	if err != nil {
		return err
	}

	res, err := clouddk.DoClientRequest(&clientSettings, "POST", fmt.Sprintf("cloudservers/%s/network-interfaces/%s/firewall-rules", serverID, networkInterfaceID), reqBody, []int{200}, 60, 10)

	if err != nil {
		resourceServerUnlock(d, m, serverID)

		return err
	}

	err = resourceServerUnlock(d, m, serverID)

	if err != nil {
		return err
	}

	firewallRule := clouddk.FirewallRuleBody{}
	err = json.NewDecoder(res.Body).Decode(&firewallRule)

	if err != nil {
		return err
	}

	return dataSourceFirewallRuleReadResponseBody(d, m, &firewallRule)
}

// resourceFirewallRuleRead reads information about an existing firewall rule.
func resourceFirewallRuleRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	firewallRuleID := d.Id()
	networkInterfaceID := d.Get(dataSourceFirewallRuleNetworkInterfaceIDKey).(string)
	serverID := d.Get(dataSourceFirewallRuleServerIDKey).(string)

	req, err := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/network-interfaces/%s/firewall-rules/%s", serverID, networkInterfaceID, firewallRuleID), new(bytes.Buffer))

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

		return fmt.Errorf("Failed to read the firewall rule information - Reason: The API responded with HTTP %s", res.Status)
	}

	firewallRule := clouddk.FirewallRuleBody{}
	err = json.NewDecoder(res.Body).Decode(&firewallRule)

	if err != nil {
		return err
	}

	return dataSourceFirewallRuleReadResponseBody(d, m, &firewallRule)
}

// resourceFirewallRuleUpdate updates an existing firewall rule.
func resourceFirewallRuleUpdate(d *schema.ResourceData, m interface{}) error {
	serverID := d.Get(dataSourceFirewallRuleServerIDKey).(string)

	clientSettings := m.(clouddk.ClientSettings)

	firewallRuleID := d.Id()
	networkInterfaceID := d.Get(dataSourceFirewallRuleNetworkInterfaceIDKey).(string)
	address := strings.Split(d.Get(dataSourceFirewallRuleAddressKey).(string), "/")

	if len(address) != 2 {
		return fmt.Errorf("Invalid address '%s' for firewall rule (must be defined as x.x.x.x/x)", d.Get(dataSourceFirewallRuleAddressKey).(string))
	}

	bits, err := strconv.Atoi(address[1])

	if err != nil {
		return fmt.Errorf("Invalid address '%s' for firewall rule (%s)", d.Get(dataSourceFirewallRuleAddressKey).(string), err.Error())
	}

	body := clouddk.FirewallRuleCreateBody{
		Command:  d.Get(dataSourceFirewallRuleCommandKey).(string),
		Protocol: d.Get(dataSourceFirewallRuleProtocolKey).(string),
		Address:  address[0],
		Bits:     clouddk.CustomInt(bits),
		Port:     d.Get(dataSourceFirewallRulePortKey).(string),
	}

	reqBody := new(bytes.Buffer)
	err = json.NewEncoder(reqBody).Encode(body)

	if err != nil {
		return err
	}

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	err = resourceServerLock(d, m, serverID)

	if err != nil {
		return err
	}

	res, err := clouddk.DoClientRequest(&clientSettings, "PUT", fmt.Sprintf("cloudservers/%s/network-interfaces/%s/firewall-rules/%s", serverID, networkInterfaceID, firewallRuleID), reqBody, []int{200}, 60, 10)

	if err != nil {
		resourceServerUnlock(d, m, serverID)

		return err
	}

	err = resourceServerUnlock(d, m, serverID)

	if err != nil {
		return err
	}

	firewallRule := clouddk.FirewallRuleBody{}
	err = json.NewDecoder(res.Body).Decode(&firewallRule)

	if err != nil {
		return err
	}

	return dataSourceFirewallRuleReadResponseBody(d, m, &firewallRule)
}

// resourceFirewallRuleDelete deletes an existing firewall rule.
func resourceFirewallRuleDelete(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	serverID := d.Get(dataSourceFirewallRuleServerIDKey).(string)
	firewallRuleID := d.Id()
	networkInterfaceID := d.Get(dataSourceFirewallRuleNetworkInterfaceIDKey).(string)

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	err := resourceServerLock(d, m, serverID)

	if err != nil {
		return err
	}

	_, err = clouddk.DoClientRequest(&clientSettings, "DELETE", fmt.Sprintf("cloudservers/%s/network-interfaces/%s/firewall-rules/%s", serverID, networkInterfaceID, firewallRuleID), new(bytes.Buffer), []int{200, 404}, 60, 10)

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
