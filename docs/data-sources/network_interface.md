---
layout: page
title: clouddk_network_interface
permalink: /data-sources/network_interface
nav_order: 7
parent: Data Sources
---

# Data Source: clouddk_network_interface

Retrieves information about a network interface for a server.

## Example Usage

```
data "clouddk_network_interface" "example" {
  id        = element(flatten(clouddk_server.example.network_interface_ids), 0)
  server_id = clouddk_server.example.id
}
```

## Argument Reference

* `id` - (Required) This is the network interface's identifier.
* `server_id` - (Required) This is the server's identifier.

## Attribute Reference

* `addresses` - This is the IP addresses assigned to the network interface.
* `default_firewall_rule` - This is the default firewall rule for the network interface.
* `firewall_rules_addresses` - This is the CIDR blocks for the firewall rules assigned to the network interface.
* `firewall_rules_commands` - This is the commands for the firewall rules assigned to the network interface.
* `firewall_rules_ids` - This is the identifiers for the firewall rules assigned to the network interface.
* `firewall_rules_ports` - This is the ports for the firewall rules assigned to the network interface.
* `firewall_rules_protocols` - This is the protocols for the firewall rules assigned to the network interface.
* `gateways` - This is the gateways assigned to the network interface.
* `label` - This is the label for the network interface.
* `netmasks` - This is the netmasks assigned to the network interface.
* `networks` - This is the networks assigned to the network interface.
* `primary` - Whether a network interface is the primary interface.
* `rate_limit` - This is the rate limit for the network interface.
