---
layout: page
title: clouddk_ip_addresses
permalink: /data-sources/ip_addresses
nav_order: 5
parent: Data Sources
---

# Data Source: clouddk_ip_addresses

Retrieves information about the IP addresses and gateways for a server.

## Example Usage

```
data "clouddk_ip_addresses" "example" {
  id = clouddk_server.example.id
}
```

## Argument Reference

* `id` - (Required) This is the server's identifier.

## Attribute Reference

* `addresses` - This is the IP addresses assigned to the server's network interfaces.
* `gateways` - This is the gateways assigned to the server's network interfaces.
* `netmasks` - This is the netmasks assigned to the server's network interfaces.
* `network_interface_ids` - This is the network interface identifiers.
* `networks` - This is the networks assigned to the server's network interfaces.
