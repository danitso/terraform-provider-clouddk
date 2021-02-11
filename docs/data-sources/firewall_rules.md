---
layout: page
title: clouddk_firewall_rules
permalink: /data-sources/firewall_rules
nav_order: 4
parent: Data Sources
---

# Data Source: clouddk_firewall_rules

Retrieves information about the firewall rules for a server.

## Example Usage

```
data "clouddk_firewall_rules" "example" {
  id        = element(flatten(clouddk_server.example.network_interface_ids), 0)
  server_id = clouddk_server.example.id
}
```

## Argument Reference

* `id` - (Required) This is the network interface's identifier.
* `server_id` - (Required) This is the server's identifier.

## Attribute Reference

* `addresses` - This is the CIDR blocks for the firewall rules assigned to the network interface.
* `commands` - This is the commands for the firewall rules assigned to the network interface.
* `ids` - This is the identifiers for the firewall rules assigned to the network interface.
* `ports` - This is the ports for the firewall rules assigned to the network interface.
* `protocols` - This is the protocols for the firewall rules assigned to the network interface.
