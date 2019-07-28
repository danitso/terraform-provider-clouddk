resource "clouddk_server" "example" {
  hostname      = "terraform-provider-clouddk-example"
  label         = "Terraform Example"
  root_password = "${random_string.clouddk_server_example_root_password.result}"

  location_id = "${element(data.clouddk_locations.example.ids, 0)}"
  package_id  = "${element(data.clouddk_packages.example.ids, index(data.clouddk_packages.example.names, "clouddk.s1"))}"
  template_id = "${element(data.clouddk_templates.example.ids, index(data.clouddk_templates.example.ids, "ubuntu-18.04-x64"))}"

  connection {
    type = "ssh"

    agent = false

    host     = "${element(flatten(self.network_interface_addresses), 0)}"
    port     = 22
    user     = "root"
    password = "${random_string.clouddk_server_example_root_password.result}"

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

resource "random_string" "clouddk_server_example_root_password" {
  length  = 32
  special = false
}

output "resource_clouddk_server_example_booted" {
  description = "Whether the server has been booted"
  value       = "${clouddk_server.example.booted}"
}

output "resource_clouddk_server_example_cpus" {
  description = "The server's CPU count"
  value       = "${clouddk_server.example.cpus}"
}

output "resource_clouddk_server_example_disk_ids" {
  description = "The server's disk identifiers"
  value       = "${clouddk_server.example.disk_ids}"
}

output "resource_clouddk_server_example_disk_labels" {
  description = "The server's disk labels"
  value       = "${clouddk_server.example.disk_labels}"
}

output "resource_clouddk_server_example_disk_primary" {
  description = "Whether the disk is the primary disk"
  value       = "${clouddk_server.example.disk_primary}"
}

output "resource_clouddk_server_example_disk_sizes" {
  description = "The server's disk identifiers"
  value       = "${clouddk_server.example.disk_sizes}"
}

output "resource_clouddk_server_example_hostname" {
  description = "The server hostname"
  value       = "${clouddk_server.example.hostname}"
}

output "resource_clouddk_server_example_id" {
  description = "The server identifier"
  value       = "${clouddk_server.example.id}"
}

output "resource_clouddk_server_example_label" {
  description = "The server label"
  value       = "${clouddk_server.example.label}"
}

output "resource_clouddk_server_example_location_id" {
  description = "The location identifier"
  value       = "${clouddk_server.example.location_id}"
}

output "resource_clouddk_server_example_location_name" {
  description = "The location name"
  value       = "${clouddk_server.example.location_name}"
}

output "resource_clouddk_server_example_memory" {
  description = "The server's memory allocation in gigabytes"
  value       = "${clouddk_server.example.memory}"
}

output "resource_clouddk_server_example_network_interface_addresses" {
  description = "The IP addresses assigned to the server's network interfaces"
  value       = "${clouddk_server.example.network_interface_addresses}"
}

output "resource_clouddk_server_example_network_interface_default_firewall_rules" {
  description = "The default firewall rules for the server's network interfaces"
  value       = "${clouddk_server.example.network_interface_default_firewall_rules}"
}

output "resource_clouddk_server_example_network_interface_firewall_rules_addresses" {
  description = "The commands for the firewall rules assigned to the server's network interfaces"
  value       = "${clouddk_server.example.network_interface_firewall_rules_addresses}"
}

output "resource_clouddk_server_example_network_interface_firewall_rules_commands" {
  description = "The commands for the firewall rules assigned to the server's network interfaces"
  value       = "${clouddk_server.example.network_interface_firewall_rules_commands}"
}

output "resource_clouddk_server_example_network_interface_firewall_rules_ids" {
  description = "The identifiers for the firewall rules assigned to the server's network interfaces"
  value       = "${clouddk_server.example.network_interface_firewall_rules_ids}"
}

output "resource_clouddk_server_example_network_interface_firewall_rules_ports" {
  description = "The ports for the firewall rules assigned to the server's network interfaces"
  value       = "${clouddk_server.example.network_interface_firewall_rules_ports}"
}

output "resource_clouddk_server_example_network_interface_firewall_rules_protocols" {
  description = "The protocols for the firewall rules assigned to the server's network interfaces"
  value       = "${clouddk_server.example.network_interface_firewall_rules_protocols}"
}

output "resource_clouddk_server_example_network_interface_gateways" {
  description = "The gateways assigned to the server's network interfaces"
  value       = "${clouddk_server.example.network_interface_gateways}"
}

output "resource_clouddk_server_example_network_interface_ids" {
  description = "The server's network interface identifiers"
  value       = "${clouddk_server.example.network_interface_ids}"
}

output "resource_clouddk_server_example_network_interface_labels" {
  description = "The server's network interface labels"
  value       = "${clouddk_server.example.network_interface_labels}"
}

output "resource_clouddk_server_example_network_interface_netmasks" {
  description = "The netmasks assigned to the server's network interfaces"
  value       = "${clouddk_server.example.network_interface_netmasks}"
}

output "resource_clouddk_server_example_network_interface_networks" {
  description = "The networks assigned to the server's network interfaces"
  value       = "${clouddk_server.example.network_interface_networks}"
}

output "resource_clouddk_server_example_network_interface_primary" {
  description = "Whether the network interface is the primary interface"
  value       = "${clouddk_server.example.network_interface_primary}"
}

output "resource_clouddk_server_example_network_interface_rate_limits" {
  description = "The rate limits for the server's network interfaces"
  value       = "${clouddk_server.example.network_interface_rate_limits}"
}

output "resource_clouddk_server_example_package_id" {
  description = "The package identifier"
  value       = "${clouddk_server.example.package_id}"
}

output "resource_clouddk_server_example_package_name" {
  description = "The package name"
  value       = "${clouddk_server.example.package_name}"
}

output "resource_clouddk_server_example_template_id" {
  description = "The template identifier"
  value       = "${clouddk_server.example.template_id}"
}

output "resource_clouddk_server_example_template_name" {
  description = "The template name"
  value       = "${clouddk_server.example.template_name}"
}
