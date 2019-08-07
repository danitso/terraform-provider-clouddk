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

const (
	resourceServerHostnameKey                                   = "hostname"
	resourceServerLabelKey                                      = "label"
	resourceServerLocationIDKey                                 = "location_id"
	resourceServerPrimaryNetworkInterfaceDefaultFirewallRuleKey = "primary_network_interface_default_firewall_rule"
	resourceServerPrimaryNetworkInterfaceLabelKey               = "primary_network_interface_label"
	resourceServerPackageIDKey                                  = "package_id"
	resourceServerRootPasswordKey                               = "root_password"
	resourceServerTemplateIDKey                                 = "template_id"
)

var (
	serverActionMutex = &sync.Mutex{}
	serverMap         = make(map[string]*sync.Mutex)
	serverMapMutex    = &sync.Mutex{}
)

// resourceServer manages a server.
func resourceServer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			resourceServerHostnameKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The hostname",
			},
			resourceServerLabelKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The label",
			},
			resourceServerLocationIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The location identifier",
				ForceNew:    true,
			},
			resourceServerPackageIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The package identifier",
			},
			resourceServerPrimaryNetworkInterfaceDefaultFirewallRuleKey: &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "ACCEPT",
				Description: "The default firewall rule for the primary network interface",
			},
			resourceServerPrimaryNetworkInterfaceLabelKey: &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Primary Network Interface",
				Description: "The label for the primary network interface",
			},
			resourceServerRootPasswordKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The root password",
				ForceNew:    true,
				Sensitive:   true,
			},
			resourceServerTemplateIDKey: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The template identifier",
				ForceNew:    true,
			},
			dataSourceServerBootedKey: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the server has been booted",
			},
			dataSourceServerCPUsKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The server's CPU count",
			},
			dataSourceServerDiskIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServerDiskLabelsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk labels",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServerDiskPrimaryKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whether a disk is the primary disk",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			dataSourceServerDiskSizesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's disk sizes in gigabytes",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			dataSourceServerLocationNameKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The location name",
			},
			dataSourceServerMemoryKey: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The server's memory allocation in megabytes",
			},
			dataSourceServerNetworkInterfaceAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The IP addresses assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceDefaultFirewallRulesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The default firewall rules for the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServerNetworkInterfaceFirewallRulesAddressesKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The CIDR blocks for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceFirewallRulesCommandsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The commands for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceFirewallRulesIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The identifiers for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceFirewallRulesPortsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The ports of the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceFirewallRulesProtocolsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The protocols for the firewall rules assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceGatewaysKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The gateways assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceIdsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's network interface identifiers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServerNetworkInterfaceLabelsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server's network interface labels",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			dataSourceServerNetworkInterfaceNetmasksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The netmasks assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfaceNetworksKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The networks assigned to the server's network interfaces",
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			dataSourceServerNetworkInterfacePrimaryKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Whether a network interface is the primary interface",
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			dataSourceServerNetworkInterfaceRateLimitsKey: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The rate limits for the server's network interfaces",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			dataSourceServerPackageNameKey: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The package name",
			},
			dataSourceServerTemplateNameKey: &schema.Schema{
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

// resourceServerCreate creates a server.
func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	body := clouddk.ServerCreateBody{
		Hostname:            d.Get(resourceServerHostnameKey).(string),
		Label:               d.Get(resourceServerLabelKey).(string),
		InitialRootPassword: d.Get(resourceServerRootPasswordKey).(string),
		Package:             d.Get(resourceServerPackageIDKey).(string),
		Template:            d.Get(resourceServerTemplateIDKey).(string),
		Location:            d.Get(resourceServerLocationIDKey).(string),
	}

	reqBody := new(bytes.Buffer)
	err := json.NewEncoder(reqBody).Encode(body)

	if err != nil {
		return err
	}

	// Due to an API issue which causes global server actions to fail, if we perform them too fast, we need to do one action at a time.
	serverActionMutex.Lock()

	res, err := clouddk.DoClientRequest(&clientSettings, "POST", "cloudservers", reqBody, []int{200}, 60, 10)

	serverActionMutex.Unlock()

	if err != nil {
		return err
	}

	server := clouddk.ServerBody{}
	err = json.NewDecoder(res.Body).Decode(&server)

	if err != nil {
		return err
	}

	err = dataSourceServerReadResponseBody(d, m, &server)

	if err != nil {
		return err
	}

	if d.Get(dataSourceServerBootedKey).(bool) {
		return nil
	}

	// Wait for the server to boot before proceeding as we may otherwise cause timeouts in provisioners.
	err = resourceServerWaitForBootFlag(d, m, &server)

	if err != nil {
		return err
	}

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	err = resourceServerLock(d, m, d.Id())

	if err != nil {
		return err
	}

	// We should now be able to change the properties for the primary network interface.
	err = resourceServerUpdatePrimaryNetworkInterface(d, m, &server)

	if err != nil {
		resourceServerUnlock(d, m, d.Id())

		return nil
	}

	// We need to release the lock for the server to allow other operations to continue.
	err = resourceServerUnlock(d, m, d.Id())

	if err != nil {
		return err
	}

	return nil
}

// resourceServerRead reads information about an existing server.
func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	req, err := clouddk.GetClientRequestObject(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s", d.Id()), new(bytes.Buffer))

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

		return fmt.Errorf("Failed to read the server information - Reason: The API responded with HTTP %s", res.Status)
	}

	server := clouddk.ServerBody{}
	err = json.NewDecoder(res.Body).Decode(&server)

	if err != nil {
		return err
	}

	err = dataSourceServerReadResponseBody(d, m, &server)

	if err != nil {
		return err
	}

	if len(server.NetworkInterfaces) > 0 {
		d.Set(resourceServerPrimaryNetworkInterfaceDefaultFirewallRuleKey, server.NetworkInterfaces[0].DefaultFirewallRule)
		d.Set(resourceServerPrimaryNetworkInterfaceLabelKey, server.NetworkInterfaces[0].Label)
	}

	return nil
}

// resourceServerUpdate updates an existing server.
func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	body := clouddk.ServerUpdateBody{
		Hostname: d.Get(resourceServerHostnameKey).(string),
		Label:    d.Get(resourceServerLabelKey).(string),
	}

	reqBody := new(bytes.Buffer)
	err := json.NewEncoder(reqBody).Encode(body)

	if err != nil {
		return err
	}

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	err = resourceServerLock(d, m, d.Id())

	if err != nil {
		return err
	}

	// We should now be able to proceed without any issues.
	res, err := clouddk.DoClientRequest(&clientSettings, "PUT", fmt.Sprintf("cloudservers/%s", d.Id()), reqBody, []int{200}, 60, 10)

	if err != nil {
		return err
	}

	server := clouddk.ServerBody{}
	err = json.NewDecoder(res.Body).Decode(&server)

	if err != nil {
		return err
	}

	// We also need to upgrade the settings for the primary network interface.
	err = resourceServerUpdatePrimaryNetworkInterface(d, m, &server)

	if err != nil {
		resourceServerUnlock(d, m, d.Id())

		return err
	}

	// In case the package has changed, we need to upgrade or downgrade the server.
	if d.HasChange(resourceServerPackageIDKey) {
		upgradeBody := clouddk.ServerUpgradeBody{
			Package:     d.Get(resourceServerPackageIDKey).(string),
			UpgradeDisk: false,
		}

		reqBody = new(bytes.Buffer)
		err = json.NewEncoder(reqBody).Encode(upgradeBody)

		if err != nil {
			resourceServerUnlock(d, m, d.Id())

			return err
		}

		res, resErr := clouddk.DoClientRequest(&clientSettings, "POST", fmt.Sprintf("cloudservers/%s/upgrade", d.Id()), reqBody, []int{200}, 60, 10)

		if resErr != nil {
			resourceServerUnlock(d, m, d.Id())

			return resErr
		}

		server = clouddk.ServerBody{}
		err = json.NewDecoder(res.Body).Decode(&server)

		if err != nil {
			resourceServerUnlock(d, m, d.Id())

			return err
		}

		err = dataSourceServerReadResponseBody(d, m, &server)

		if err != nil {
			resourceServerUnlock(d, m, d.Id())

			return err
		}
	}

	// We need to release the lock for the server to allow other operations to continue.
	err = resourceServerUnlock(d, m, d.Id())

	if err != nil {
		return err
	}

	return nil
}

// resourceServerUpdatePrimaryNetworkInterface updates the primary interface on an existing server.
func resourceServerUpdatePrimaryNetworkInterface(d *schema.ResourceData, m interface{}, server *clouddk.ServerBody) error {
	clientSettings := m.(clouddk.ClientSettings)

	networkInterfaceUpdateBody := clouddk.NetworkInterfaceUpdateBody{
		DefaultFirewallRule: d.Get(resourceServerPrimaryNetworkInterfaceDefaultFirewallRuleKey).(string),
		Label:               d.Get(resourceServerPrimaryNetworkInterfaceLabelKey).(string),
	}

	reqBody := new(bytes.Buffer)
	err := json.NewEncoder(reqBody).Encode(networkInterfaceUpdateBody)

	if err != nil {
		return err
	}

	res, err := clouddk.DoClientRequest(&clientSettings, "PUT", fmt.Sprintf("cloudservers/%s/network-interfaces/%s", server.Identifier, server.NetworkInterfaces[0].Identifier), reqBody, []int{200}, 60, 10)

	if err != nil {
		return err
	}

	networkInterfaceBody := clouddk.NetworkInterfaceBody{}
	err = json.NewDecoder(res.Body).Decode(&networkInterfaceBody)

	if err != nil {
		return err
	}

	server.NetworkInterfaces[0] = networkInterfaceBody

	return dataSourceServerReadResponseBody(d, m, server)
}

// resourceServerDelete deletes an existing server.
func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)

	// We need to acquire the lock for the server to reduce the risk of race conditions.
	err := resourceServerLock(d, m, d.Id())

	if err != nil {
		return err
	}

	// We should now be able to proceed without any issues.
	_, err = clouddk.DoClientRequest(&clientSettings, "DELETE", fmt.Sprintf("cloudservers/%s", d.Id()), new(bytes.Buffer), []int{200, 404}, 60, 10)

	if err != nil {
		return err
	}

	// We need to release the lock for the server to allow other operations to continue.
	err = resourceServerUnlock(d, m, d.Id())

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

// resourceServerLock acquires the lock for a specific server.
func resourceServerLock(d *schema.ResourceData, m interface{}, serverID string) error {
	clientSettings := m.(clouddk.ClientSettings)

	retryLimit := 90
	retryDelay := 10

	// Acquire the lock for the serverMap variable.
	log.Printf("[DEBUG] Acquiring lock for server map (id: %s)", serverID)
	serverMapMutex.Lock()

	// Create a mutex for the specified server, if none already exists.
	if serverMap[serverID] == nil {
		log.Printf("[DEBUG] Creating mutex for server (id: %s)", serverID)
		serverMap[serverID] = &sync.Mutex{}
	}

	// Now that we know a mutex for the server exists, we can unlock the mutex for the map and acquire the lock for the server instead.
	log.Printf("[DEBUG] Releasing lock for server map (id: %s)", serverID)
	serverMapMutex.Unlock()

	log.Printf("[DEBUG] Acquiring lock for server (id: %s)", serverID)
	serverMap[serverID].Lock()

	// We can now go ahead and retrieve the transactions for the server. We will keep doing this until all transactions are eithed failed or completed.
	timeDelay := int64(retryDelay)
	timeMax := float64(retryLimit * retryDelay)
	timeStart := time.Now()
	timeElapsed := timeStart.Sub(timeStart)

	continueToWait := false

	for timeElapsed.Seconds() < timeMax {
		if int64(timeElapsed.Seconds())%timeDelay == 0 {
			res, err := clouddk.DoClientRequest(&clientSettings, "GET", fmt.Sprintf("cloudservers/%s/logs", serverID), new(bytes.Buffer), []int{200}, 1, 1)

			if err != nil {
				return err
			}

			logsList := clouddk.LogsListBody{}
			err = json.NewDecoder(res.Body).Decode(&logsList)

			if err != nil {
				return err
			}

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
		log.Printf("[DEBUG] Releasing lock for server (id: %s)", serverID)
		serverMap[serverID].Unlock()

		return fmt.Errorf("Timeout while waiting for transactions to end (id: %s)", serverID)
	}

	return nil
}

// resourceServerUnlock releases the lock for a specific server.
func resourceServerUnlock(d *schema.ResourceData, m interface{}, serverID string) error {
	if serverMap[serverID] == nil {
		return fmt.Errorf("Cannot unlock a server which has never been locked during this session (id: %s)", serverID)
	}

	log.Printf("[DEBUG] Releasing lock for server (id: %s)", serverID)
	serverMap[serverID].Unlock()

	return nil
}

// resourceServerWaitForBootFlag waits for the boot flag to be toggled.
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

			err = json.NewDecoder(res.Body).Decode(server)

			if err != nil {
				return err
			}

			if server.Booted {
				return dataSourceServerReadResponseBody(d, m, server)
			}
		}

		time.Sleep(200 * time.Millisecond)

		timeElapsed = time.Now().Sub(timeStart)
	}

	return fmt.Errorf("The server '%s' (id: %s) does not appear to be able to boot", d.Get(resourceServerHostnameKey).(string), server.Identifier)
}
