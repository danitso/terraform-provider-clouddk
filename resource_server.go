package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform/helper/schema"
)

const ResourceServerHostnameKey = "hostname"
const ResourceServerLabelKey = "label"
const ResourceServerLocationIdKey = "location_id"
const ResourceServerPrimaryNetworkInterfaceDefaultFirewallRuleKey = "primary_network_interface_default_firewall_rule"
const ResourceServerPrimaryNetworkInterfaceLabelKey = "primary_network_interface_label"
const ResourceServerPackageIdKey = "package_id"
const ResourceServerRootPasswordKey = "root_password"
const ResourceServerTemplateIdKey = "template_id"

var serverMap = make(map[string]*sync.Mutex)
var serverMapMutex = &sync.Mutex{}

// resourceServer() manages a server.
func resourceServer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			ResourceServerHostnameKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The hostname",
			},
			ResourceServerLabelKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The label",
			},
			ResourceServerLocationIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The location identifier",
				ForceNew:    true,
			},
			ResourceServerPackageIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The package identifier",
				ForceNew:    true,
			},
			ResourceServerPrimaryNetworkInterfaceDefaultFirewallRuleKey: &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "ACCEPT",
				Description: "The default firewall rule for the primary network interface",
			},
			ResourceServerPrimaryNetworkInterfaceLabelKey: &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Primary Network Interface",
				Description: "The label for the primary network interface",
			},
			ResourceServerRootPasswordKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The root password",
				ForceNew:    true,
				Sensitive:   true,
			},
			ResourceServerTemplateIdKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The template identifier",
				ForceNew:    true,
			},
			DataSourceServerBootedKey: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the server has been booted",
			},
			DataSourceServerCPUsKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The server's CPU count",
			},
			DataSourceServerDiskIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServerDiskLabelsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk labels",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServerDiskPrimaryKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whether a disk is the primary disk",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			DataSourceServerDiskSizesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk sizes in gigabytes",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			DataSourceServerLocationNameKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The location name",
			},
			DataSourceServerMemoryKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The server's memory allocation in megabytes",
			},
			DataSourceServerNetworkInterfaceAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The IP addresses assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceDefaultFirewallRulesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The default firewall rules for the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServerNetworkInterfaceFirewallRulesAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The CIDR blocks for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceFirewallRulesCommandsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The commands for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceFirewallRulesIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The identifiers for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceFirewallRulesPortsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The ports of the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceFirewallRulesProtocolsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The protocols for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceGatewaysKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The gateways assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's network interface identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServerNetworkInterfaceLabelsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's network interface labels",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServerNetworkInterfaceNetmasksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The netmasks assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfaceNetworksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The networks assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			DataSourceServerNetworkInterfacePrimaryKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whether a network interface is the primary interface",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			DataSourceServerNetworkInterfaceRateLimitsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The rate limits for the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			DataSourceServerPackageNameKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The package name",
			},
			DataSourceServerTemplateNameKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The template name",
			},
		},

		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,
	}
}

// resourceServerCreate() creates a server.
func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	body := clouddk.ServerCreateBody{
		Hostname:            d.Get(ResourceServerHostnameKey).(string),
		Label:               d.Get(ResourceServerLabelKey).(string),
		InitialRootPassword: d.Get(ResourceServerRootPasswordKey).(string),
		Package:             d.Get(ResourceServerPackageIdKey).(string),
		Template:            d.Get(ResourceServerTemplateIdKey).(string),
		Location:            d.Get(ResourceServerLocationIdKey).(string),
	}

	reqBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(reqBody).Encode(body)

	if encodeErr != nil {
		return encodeErr
	}

	res, resErr := clouddk.DoClientRequest(&clientSettings, "POST", "cloudservers", reqBody, []int{200}, 60, 10)

	if resErr != nil {
		return resErr
	}

	server := clouddk.ServerBody{}
	json.NewDecoder(res.Body).Decode(&server)

	parseErr := dataSourceServerReadResponseBody(d, m, &server)

	if parseErr != nil {
		return parseErr
	}

	if d.Get(DataSourceServerBootedKey).(bool) {
		return nil
	}

	// Wait for the server to boot before proceeding as we may otherwise cause timeouts in provisioners.
	bootErr := resourceServerWaitForBootFlag(d, m, &server)

	if bootErr != nil {
		return bootErr
	}

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	lockErr := resourceServerLock(d, m, d.Id())

	if lockErr != nil {
		return lockErr
	}

	// We should now be able to change the properties for the primary network interface.
	updateNetworkError := resourceServerUpdatePrimaryNetworkInterface(d, m, &server)

	if updateNetworkError != nil {
		resourceServerUnlock(d, m, d.Id())

		return nil
	}

	// We need to release the lock for the server to allow other operations to continue.
	lockErr = resourceServerUnlock(d, m, d.Id())

	if lockErr != nil {
		return lockErr
	}

	return nil
}

// resourceServerRead reads information about an existing server.
func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	req, reqErr := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s", d.Id()), new(bytes.Buffer))

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

		return fmt.Errorf("Failed to read the server information - Reason: The API responded with HTTP %s", res.Status)
	}

	server := clouddk.ServerBody{}
	json.NewDecoder(res.Body).Decode(&server)

	parseErr := dataSourceServerReadResponseBody(d, m, &server)

	if parseErr != nil {
		return parseErr
	}

	d.Set(ResourceServerPrimaryNetworkInterfaceDefaultFirewallRuleKey, server.NetworkInterfaces[0].DefaultFirewallRule)
	d.Set(ResourceServerPrimaryNetworkInterfaceLabelKey, server.NetworkInterfaces[0].Label)

	return nil
}

// resourceServerUpdate updates an existing server.
func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	body := clouddk.ServerUpdateBody{
		Hostname: d.Get(ResourceServerHostnameKey).(string),
		Label:    d.Get(ResourceServerLabelKey).(string),
	}

	reqBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(reqBody).Encode(body)

	if encodeErr != nil {
		return encodeErr
	}

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	lockErr := resourceServerLock(d, m, d.Id())

	if lockErr != nil {
		return lockErr
	}

	// We should now be able to proceed without any issues.
	res, resErr := clouddk.DoClientRequest(&clientSettings, "PUT", fmt.Sprintf("cloudservers/%s", d.Id()), reqBody, []int{200}, 60, 10)

	if resErr != nil {
		return resErr
	}

	server := clouddk.ServerBody{}
	json.NewDecoder(res.Body).Decode(&server)

	updateNetworkError := resourceServerUpdatePrimaryNetworkInterface(d, m, &server)

	if updateNetworkError != nil {
		resourceServerUnlock(d, m, d.Id())

		return nil
	}

	// We need to release the lock for the server to allow other operations to continue.
	lockErr = resourceServerUnlock(d, m, d.Id())

	if lockErr != nil {
		return lockErr
	}

	return nil
}

// resourceServerUpdatePrimaryNetworkInterface updates the primary interface on an existing server.
func resourceServerUpdatePrimaryNetworkInterface(d *schema.ResourceData, m interface{}, server *clouddk.ServerBody) error {
	clientSettings := m.(clouddk.ClientSettings)

	networkInterfaceUpdateBody := clouddk.NetworkInterfaceUpdateBody{
		DefaultFirewallRule: d.Get(ResourceServerPrimaryNetworkInterfaceDefaultFirewallRuleKey).(string),
		Label:               d.Get(ResourceServerPrimaryNetworkInterfaceLabelKey).(string),
	}

	reqBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(reqBody).Encode(networkInterfaceUpdateBody)

	if encodeErr != nil {
		return encodeErr
	}

	res, resErr := clouddk.DoClientRequest(&clientSettings, "PUT", fmt.Sprintf("cloudservers/%s/network-interfaces/%s", server.Identifier, server.NetworkInterfaces[0].Identifier), reqBody, []int{200}, 60, 10)

	if resErr != nil {
		return resErr
	}

	networkInterfaceBody := clouddk.NetworkInterfaceBody{}
	json.NewDecoder(res.Body).Decode(&networkInterfaceBody)

	server.NetworkInterfaces[0] = networkInterfaceBody

	return dataSourceServerReadResponseBody(d, m, server)
}

// resourceServerDelete deletes an existing server.
func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	lockErr := resourceServerLock(d, m, d.Id())

	if lockErr != nil {
		return lockErr
	}

	// We should now be able to proceed without any issues.
	_, err := clouddk.DoClientRequest(&clientSettings, "DELETE", fmt.Sprintf("cloudservers/%s", d.Id()), new(bytes.Buffer), []int{200, 404}, 60, 10)

	if err != nil {
		return err
	}

	// We need to release the lock for the server to allow other operations to continue.
	lockErr = resourceServerUnlock(d, m, d.Id())

	if lockErr != nil {
		return lockErr
	}

	d.SetId("")

	return nil
}

// resourceServerLock() acquires the lock for a specific server.
func resourceServerLock(d *schema.ResourceData, m interface{}, serverId string) error {
	clientSettings := m.(clouddk.ClientSettings)

	retryLimit := 90
	retryDelay := 10

	// Acquire the lock for the serverMap variable.
	log.Printf("[DEBUG] Acquiring lock for server map (id: %s)", serverId)
	serverMapMutex.Lock()

	// Create a mutex for the specified server, if none already exists.
	if serverMap[serverId] == nil {
		log.Printf("[DEBUG] Creating mutex for server (id: %s)", serverId)
		serverMap[serverId] = &sync.Mutex{}
	}

	// Now that we know a mutex for the server exists, we can unlock the mutex for the map and acquire the lock for the server instead.
	log.Printf("[DEBUG] Releasing lock for server map (id: %s)", serverId)
	serverMapMutex.Unlock()

	log.Printf("[DEBUG] Acquiring lock for server (id: %s)", serverId)
	serverMap[serverId].Lock()

	// We can now go ahead and retrieve the transactions for the server. We will keep doing this until all transactions are eithed failed or completed.
	timeDelay := int64(retryDelay)
	timeMax := float64(retryLimit * retryDelay)
	timeStart := time.Now()
	timeElapsed := timeStart.Sub(timeStart)

	continueToWait := false

	for timeElapsed.Seconds() < timeMax {
		if int64(timeElapsed.Seconds())%timeDelay == 0 {
			res, err := clouddk.DoClientRequest(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/logs", serverId), new(bytes.Buffer), []int{200}, 1, 1)

			if err != nil {
				return err
			}

			logsList := clouddk.LogsListBody{}
			json.NewDecoder(res.Body).Decode(&logsList)

			continueToWait = false

			for _, v := range logsList {
				if v.Status == "pending" || v.Status == "running" {
					continueToWait = true

					break
				}
			}

			if !continueToWait {
				break
			}

			time.Sleep(1 * time.Second)
		}

		time.Sleep(200 * time.Millisecond)

		timeElapsed = time.Now().Sub(timeStart)
	}

	// Throw an error in case there are still transactions pending or running
	if continueToWait {
		log.Printf("[DEBUG] Releasing lock for server (id: %s)", serverId)
		serverMap[serverId].Unlock()

		return fmt.Errorf("Timeout while waiting for transactions to end (id: %s)", serverId)
	}

	return nil
}

// resourceServerUnlock() releases the lock for a specific server.
func resourceServerUnlock(d *schema.ResourceData, m interface{}, serverId string) error {
	if serverMap[serverId] == nil {
		return fmt.Errorf("Cannot unlock a server which has never been locked during this session (id: %s)", serverId)
	}

	log.Printf("[DEBUG] Releasing lock for server (id: %s)", serverId)
	serverMap[serverId].Unlock()

	return nil
}

// resourceServerWaitForBootFlag() waits for the boot flag to be toggled.
func resourceServerWaitForBootFlag(d *schema.ResourceData, m interface{}, server *clouddk.ServerBody) error {
	clientSettings := m.(clouddk.ClientSettings)

	// For some reason the API is still indicating that the server has not been booted. Let's wait a while for that to change.
	timeDelay := int64(10)
	timeMax := float64(600)
	timeStart := time.Now()
	timeElapsed := timeStart.Sub(timeStart)

	for timeElapsed.Seconds() < timeMax {
		if int64(timeElapsed.Seconds())%timeDelay == 0 {
			res, err := clouddk.DoClientRequest(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s", server.Identifier), new(bytes.Buffer), []int{200}, 1, 1)

			if err != nil {
				return err
			}

			json.NewDecoder(res.Body).Decode(server)

			if server.Booted {
				return dataSourceServerReadResponseBody(d, m, server)
			}
		}

		time.Sleep(200 * time.Millisecond)

		timeElapsed = time.Now().Sub(timeStart)
	}

	return fmt.Errorf("The server '%s' (id: %s) does not appear to be able to boot", d.Get(ResourceServerHostnameKey).(string), server.Identifier)
}
