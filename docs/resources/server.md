---
layout: page
title: clouddk_server
permalink: /resources/server
nav_order: 4
parent: Resources
---

# Resource: clouddk_server

Manages a server.

## Example Usage

```
resource "clouddk_server" "example" {
  hostname      = "terraform-provider-clouddk-example"
  label         = "Terraform Example"
  root_password = random_password.clouddk_server_example_root_password.result

  location_id = element(data.clouddk_locations.example.ids, 0)
  package_id  = element(data.clouddk_packages.example.ids, index(data.clouddk_packages.example.names, "clouddk.s1"))
  template_id = "ubuntu-18.04-x64"

  connection {
    type = "ssh"

    agent = false

    host     = element(flatten(self.network_interface_addresses), 0)
    port     = 22
    user     = "root"
    password = random_password.clouddk_server_example_root_password.result

    timeout = "300s"
  }

  provisioner "remote-exec" {
    inline = [
      "echo The server was successfully provisioned!",
      "echo Continuing in 5 seconds",
      "sleep 5",
    ]
  }
}
```

## Argument Reference

* `hostname` - (Required) This is the server's hostname.
* `label` - (Required) This is the server's label.
* `location_id` - (Required) This is the server's location.
* `package_id` - (Required) This is the server's package.
* `primary_network_interface_default_firewall_rule` - (Optional) This is the default firewall rule for the server's primary network interface.
* `primary_network_interface_label` - (Optional) This is the label for the server's primary network interface.
* `root_password` - (Required) This is the initial root password.
* `template_id` - (Required) This is the server's template.

## Attribute Reference

* `booted` - Whether the server has been booted.
* `cpus` - This is the server's CPU count.
* `disk_ids` - This is the server's disk identifiers.
* `disk_labels` - This is the server's disk labels.
* `disk_primary` - Whether a disk is the primary disk.
* `disk_sizes` - This is the server's disk sizes in gigabytes.
* `hostname` - This is the server's hostname.
* `id` - This is the server's identifier.
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
