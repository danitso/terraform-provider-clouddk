package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

// resourceFirewallRule() manages a firewall rule.
func resourceFirewallRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceFirewallRuleAddressKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The CIDR block for the firewall rule",
			},
			DataSourceFirewallRuleCommandKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The command for the firewall rule",
			},
			DataSourceFirewallRuleNetworkInterfaceIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The network interface identifier",
				ForceNew:    true,
			},
			DataSourceFirewallRulePortKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The port for the firewall rule",
			},
			DataSourceFirewallRuleProtocolKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The protocol for the firewall rule",
			},
			DataSourceFirewallRuleServerIdKey: &schema.Schema{
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

// resourceFirewallRuleCreate() creates a firewall rule.
func resourceFirewallRuleCreate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	networkInterfaceId := d.Get(DataSourceFirewallRuleNetworkInterfaceIdKey).(string)
	serverId := d.Get(DataSourceFirewallRuleServerIdKey).(string)

	address := strings.Split(d.Get(DataSourceFirewallRuleAddressKey).(string), "/")

	if len(address) != 2 {
		return fmt.Errorf("Invalid address '%s' for firewall rule (must be defined as x.x.x.x/x)", d.Get(DataSourceFirewallRuleAddressKey).(string))
	}

	bits, bitsErr := strconv.Atoi(address[1])

	if bitsErr != nil {
		return fmt.Errorf("Invalid address '%s' for firewall rule (%s)", d.Get(DataSourceFirewallRuleAddressKey).(string), bitsErr.Error())
	}

	body := FirewallRuleCreateBody{
		Command:  d.Get(DataSourceFirewallRuleCommandKey).(string),
		Protocol: d.Get(DataSourceFirewallRuleProtocolKey).(string),
		Address:  address[0],
		Bits:     bits,
		Port:     d.Get(DataSourceFirewallRulePortKey).(string),
	}

	reqBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(reqBody).Encode(body)

	if encodeErr != nil {
		return encodeErr
	}

	res, resErr := doClientRequest(&clientSettings, "POST", fmt.Sprintf("cloudservers/%s/network-interfaces/%s/firewall-rules", serverId, networkInterfaceId), reqBody, []int{200}, 60, 10)

	if resErr != nil {
		return resErr
	}

	firewallRule := FirewallRuleBody{}
	json.NewDecoder(res.Body).Decode(&firewallRule)

	return dataSourceFirewallRuleReadResponseBody(d, m, &firewallRule)
}

// resourceFirewallRuleRead reads information about an existing firewall rule.
func resourceFirewallRuleRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	firewallRuleId := d.Id()
	networkInterfaceId := d.Get(DataSourceFirewallRuleNetworkInterfaceIdKey).(string)
	serverId := d.Get(DataSourceFirewallRuleServerIdKey).(string)

	req, reqErr := getClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/network-interfaces/%s/firewall-rules/%s", serverId, networkInterfaceId, firewallRuleId), new(bytes.Buffer))

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

		return fmt.Errorf("Failed to read the firewall rule information - Reason: The API responded with HTTP %s", res.Status)
	}

	firewallRule := FirewallRuleBody{}
	json.NewDecoder(res.Body).Decode(&firewallRule)

	return dataSourceFirewallRuleReadResponseBody(d, m, &firewallRule)
}

// resourceFirewallRuleUpdate updates an existing firewall rule.
func resourceFirewallRuleUpdate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	firewallRuleId := d.Id()
	networkInterfaceId := d.Get(DataSourceFirewallRuleNetworkInterfaceIdKey).(string)
	serverId := d.Get(DataSourceFirewallRuleServerIdKey).(string)

	address := strings.Split(d.Get(DataSourceFirewallRuleAddressKey).(string), "/")

	if len(address) != 2 {
		return fmt.Errorf("Invalid address '%s' for firewall rule (must be defined as x.x.x.x/x)", d.Get(DataSourceFirewallRuleAddressKey).(string))
	}

	bits, bitsErr := strconv.Atoi(address[1])

	if bitsErr != nil {
		return fmt.Errorf("Invalid address '%s' for firewall rule (%s)", d.Get(DataSourceFirewallRuleAddressKey).(string), bitsErr.Error())
	}

	body := FirewallRuleCreateBody{
		Command:  d.Get(DataSourceFirewallRuleCommandKey).(string),
		Protocol: d.Get(DataSourceFirewallRuleProtocolKey).(string),
		Address:  address[0],
		Bits:     bits,
		Port:     d.Get(DataSourceFirewallRulePortKey).(string),
	}

	reqBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(reqBody).Encode(body)

	if encodeErr != nil {
		return encodeErr
	}

	res, resErr := doClientRequest(&clientSettings, "PUT", fmt.Sprintf("cloudservers/%s/network-interfaces/%s/firewall-rules/%s", serverId, networkInterfaceId, firewallRuleId), reqBody, []int{200}, 60, 10)

	if resErr != nil {
		return resErr
	}

	firewallRule := FirewallRuleBody{}
	json.NewDecoder(res.Body).Decode(&firewallRule)

	return dataSourceFirewallRuleReadResponseBody(d, m, &firewallRule)
}

// resourceFirewallRuleDelete deletes an existing firewall rule.
func resourceFirewallRuleDelete(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(ClientSettings)

	firewallRuleId := d.Id()
	networkInterfaceId := d.Get(DataSourceFirewallRuleNetworkInterfaceIdKey).(string)
	serverId := d.Get(DataSourceFirewallRuleServerIdKey).(string)

	_, err := doClientRequest(&clientSettings, "DELETE", fmt.Sprintf("cloudservers/%s/network-interfaces/%s/firewall-rules/%s", serverId, networkInterfaceId, firewallRuleId), new(bytes.Buffer), []int{200, 404}, 60, 10)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
