---
layout: page
title: clouddk_ip_address
permalink: /resources/ip_address
nav_order: 3
parent: Resources
---

# Resource: clouddk_ip_address

Manages a public IP address for a server.

## Example Usage

```
resource "clouddk_ip_address" "example" {
  server_id = clouddk_server.example.id
}
```

## Argument Reference

* `server_id` - (Required) This is the server's identifier.

## Attribute Reference

* `address` - This is the IP address.
* `gateway` - This is the gateway address.
* `id` - This is the IP address' identifier.
* `netmask` - This is the netmask.
* `network` - This is the network address.
* `network_interface_id` - This is the identifier for the network interface that the IP address is assigned to.
