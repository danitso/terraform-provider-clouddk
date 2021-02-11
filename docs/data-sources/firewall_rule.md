---
layout: page
title: clouddk_firewall_rule
permalink: /data-sources/firewall_rule
nav_order: 3
parent: Data Sources
---

# Data Source: clouddk_firewall_rule

Retrieves information about a firewall rule for a server.

## Example Usage

```
data "clouddk_firewall_rule" "example" {
  id                   = clouddk_firewall_rule.example.id
  network_interface_id = clouddk_firewall_rule.example.network_interface_id
  server_id            = clouddk_server.example.id
}
```

## Argument Reference

* `id` - (Required) This is the firewall rule's identifier.
* `network_interface_id` - (Required) This is the network interface's identifier.
* `server_id` - (Required) This is the server's identifier.

## Attribute Reference

* `address` - This is the CIDR block for the firewall rule.
* `command` - This is the command for the firewall rule.
* `port` - This is the port for the firewall rule.
* `protocol` - This is the protocol for the firewall rule.
