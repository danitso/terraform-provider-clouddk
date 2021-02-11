---
layout: page
title: clouddk_server
permalink: /data-sources/server
nav_order: 10
parent: Data Sources
---

# Data Source: clouddk_server

Retrieves information about a server.

## Example Usage

```
data "clouddk_server" "example" {
  id = clouddk_server.example.id
}
```

## Argument Reference

* `id` - (Required) This is the server's identifier.

## Attribute Reference

* `booted` - Whether the server has been booted.
* `cpus` - This is the server's CPU count.
* `disk_ids` - This is the server's disk identifiers.
* `disk_labels` - This is the server's disk labels.
* `disk_primary` - Whether a disk is the primary disk.
* `disk_sizes` - This is the server's disk sizes in gigabytes.
* `hostname` - This is the server's hostname.
* `label` - This is the server's label.
* `location_id` - This is the location identifier.
* `location_name` - This is the location name.
* `memory` - This is the server's memory allocation in megabytes.
* `network_interface_addresses` - This is the IP addresses assigned to the server's network interfaces.
* `network_interface_default_firewall_rules` - This is the default firewall rules for the server's network interfaces.
* `network_interface_firewall_rules_addresses` - This is the CIDR blocks for the firewall rules assigned to the server's network interfaces.
* `network_interface_firewall_rules_commands` - This is the commands for the firewall rules assigned to the server's network interfaces.
* `network_interface_firewall_rules_ids` - This is the identifiers for the firewall rules assigned to the server's network interfaces.
* `network_interface_firewall_rules_ports` - This is the ports for the firewall rules assigned to the server's network interfaces.
* `network_interface_firewall_rules_protocols` - This is the protocols for the firewall rules assigned to the server's network interfaces.
* `network_interface_gateways` - This is the gateways assigned to the server's network interfaces.
* `network_interface_ids` - This is the server's network interface identifiers.
* `network_interface_labels` - This is the server's network interface labels.
* `network_interface_netmasks` - This is the netmasks assigned to the server's network interfaces.
* `network_interface_networks` - This is the networks assigned to the server's network interfaces.
* `network_interface_primary` - Whether a network interface is the primary interface.
* `network_interface_rate_limits` - This is the rate limits for the server's network interfaces.
* `package_id` - This is the package identifier.
* `package_name` - This is the package name.
* `template_id` - This is the template identifier.
* `template_name` - This is the template name.
