---
layout: page
title: clouddk_firewall_rule
permalink: /resources/firewall_rule
nav_order: 2
parent: Resources
---

# Resource: clouddk_firewall_rule

Manages a firewall rule for a server.

## Example Usage

```
resource "clouddk_firewall_rule" "example" {
  server_id            = clouddk_server.example.id
  network_interface_id = element(flatten(clouddk_server.example.network_interface_ids), 0)

  command  = "ACCEPT"
  protocol = "TCP"
  address  = "8.8.8.8/32"
  port     = 8080
}
```

## Argument Reference

* `address` - (Required) This is the CIDR block for the firewall rule.
* `command` - (Required) This is the command for the firewall rule.
* `network_interface_id` - (Required) This is the network interface's identifier.
* `port` - (Required) This is the port for the firewall rule.
* `protocol` - (Required) This is the protocol for the firewall rule.
* `server_id` - (Required) This is the server's identifier.

## Attribute Reference

* `id` - This is the firewall rule's identifier.
