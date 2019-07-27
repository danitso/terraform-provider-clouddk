data "clouddk_server" "example" {
  id = "${element(data.clouddk_servers.example.ids, 0)}"
}

output "data_clouddk_server_example_booted" {
  description = "Whether the server has been booted"
  value       = "${data.clouddk_server.example.booted}"
}

output "data_clouddk_server_example_cpus" {
  description = "The server's CPU count"
  value       = "${data.clouddk_server.example.cpus}"
}

output "data_clouddk_server_example_disk_ids" {
  description = "The server's disk identifiers"
  value       = "${data.clouddk_server.example.disk_ids}"
}

output "data_clouddk_server_example_disk_labels" {
  description = "The server's disk labels"
  value       = "${data.clouddk_server.example.disk_labels}"
}

output "data_clouddk_server_example_disk_primary" {
  description = "Whether the disk is the primary disk"
  value       = "${data.clouddk_server.example.disk_primary}"
}

output "data_clouddk_server_example_disk_sizes" {
  description = "The server's disk identifiers"
  value       = "${data.clouddk_server.example.disk_sizes}"
}

output "data_clouddk_server_example_hostname" {
  description = "The server hostname"
  value       = "${data.clouddk_server.example.hostname}"
}

output "data_clouddk_server_example_id" {
  description = "The server identifier"
  value       = "${data.clouddk_server.example.id}"
}

output "data_clouddk_server_example_label" {
  description = "The server label"
  value       = "${data.clouddk_server.example.label}"
}

output "data_clouddk_server_example_location_id" {
  description = "The location identifier"
  value       = "${data.clouddk_server.example.location_id}"
}

output "data_clouddk_server_example_location_name" {
  description = "The location name"
  value       = "${data.clouddk_server.example.location_name}"
}

output "data_clouddk_server_example_memory" {
  description = "The server's memory allocation in gigabytes"
  value       = "${data.clouddk_server.example.memory}"
}

output "data_clouddk_server_example_network_interface_addresses" {
  description = "The IP addresses assigned to the server's network interfaces"
  value       = "${data.clouddk_server.example.network_interface_addresses}"
}

output "data_clouddk_server_example_network_interface_default_firewall_rules" {
  description = "The default firewall rules for the server's network interfaces"
  value       = "${data.clouddk_server.example.network_interface_default_firewall_rules}"
}

output "data_clouddk_server_example_network_interface_firewall_rules_addresses" {
  description = "The commands for the firewall rules assigned to the server's network interfaces"
  value       = "${data.clouddk_server.example.network_interface_firewall_rules_addresses}"
}

output "data_clouddk_server_example_network_interface_firewall_rules_commands" {
  description = "The commands for the firewall rules assigned to the server's network interfaces"
  value       = "${data.clouddk_server.example.network_interface_firewall_rules_commands}"
}

output "data_clouddk_server_example_network_interface_firewall_rules_ids" {
  description = "The identifiers for the firewall rules assigned to the server's network interfaces"
  value       = "${data.clouddk_server.example.network_interface_firewall_rules_ids}"
}

output "data_clouddk_server_example_network_interface_firewall_rules_ports" {
  description = "The ports for the firewall rules assigned to the server's network interfaces"
  value       = "${data.clouddk_server.example.network_interface_firewall_rules_ports}"
}

output "data_clouddk_server_example_network_interface_firewall_rules_protocols" {
  description = "The protocols for the firewall rules assigned to the server's network interfaces"
  value       = "${data.clouddk_server.example.network_interface_firewall_rules_protocols}"
}

output "data_clouddk_server_example_network_interface_gateways" {
  description = "The gateways assigned to the server's network interfaces"
  value       = "${data.clouddk_server.example.network_interface_gateways}"
}

output "data_clouddk_server_example_network_interface_ids" {
  description = "The server's network interface identifiers"
  value       = "${data.clouddk_server.example.network_interface_ids}"
}

output "data_clouddk_server_example_network_interface_labels" {
  description = "The server's network interface labels"
  value       = "${data.clouddk_server.example.network_interface_labels}"
}

output "data_clouddk_server_example_network_interface_netmasks" {
  description = "The netmasks assigned to the server's network interfaces"
  value       = "${data.clouddk_server.example.network_interface_netmasks}"
}

output "data_clouddk_server_example_network_interface_networks" {
  description = "The networks assigned to the server's network interfaces"
  value       = "${data.clouddk_server.example.network_interface_networks}"
}

output "data_clouddk_server_example_network_interface_primary" {
  description = "Whether the network interface is the primary interface"
  value       = "${data.clouddk_server.example.network_interface_primary}"
}

output "data_clouddk_server_example_network_interface_rate_limits" {
  description = "The rate limits for the server's network interfaces"
  value       = "${data.clouddk_server.example.network_interface_rate_limits}"
}

output "data_clouddk_server_example_package_id" {
  description = "The package identifier"
  value       = "${data.clouddk_server.example.package_id}"
}

output "data_clouddk_server_example_package_name" {
  description = "The package name"
  value       = "${data.clouddk_server.example.package_name}"
}

output "data_clouddk_server_example_template_id" {
  description = "The template identifier"
  value       = "${data.clouddk_server.example.template_id}"
}

output "data_clouddk_server_example_template_name" {
  description = "The template name"
  value       = "${data.clouddk_server.example.template_name}"
}
