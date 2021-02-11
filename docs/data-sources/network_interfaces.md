---
layout: page
title: clouddk_network_interfaces
permalink: /data-sources/network_interfaces
nav_order: 8
parent: Data Sources
---

# Data Source: clouddk_network_interfaces

Retrieves information about the network interfaces for a server.

## Example Usage

```
data "clouddk_network_interfaces" "example" {
  id = clouddk_server.example.id
}
```

## Argument Reference

* `id` - (Required) This is the server's identifier.

## Attribute Reference

* `addresses` - This is the IP addresses assigned to the server's network interfaces.
* `default_firewall_rules` - This is the default firewall rules for the server's network interfaces.
* `firewall_rules_addresses` - This is the CIDR blocks for the firewall rules assigned to the server's network interfaces.
* `firewall_rules_commands` - This is the commands for the firewall rules assigned to the server's network interfaces.
* `firewall_rules_ids` - This is the identifiers for the firewall rules assigned to the server's network interfaces.
* `firewall_rules_ports` - This is the ports for the firewall rules assigned to the server's network interfaces.
* `firewall_rules_protocols` - This is the protocols for the firewall rules assigned to the server's network interfaces.
* `gateways` - This is the gateways assigned to the server's network interfaces.
* `ids` - This is the server's network interface identifiers.
* `labels` - This is the server's network interface labels.
* `netmasks` - This is the netmasks assigned to the server's network interfaces.
* `networks` - This is the networks assigned to the server's network interfaces.
* `primary` - Whether a network interface is the primary interface.
* `rate_limits` - This is the rate limits for the server's network interfaces.
