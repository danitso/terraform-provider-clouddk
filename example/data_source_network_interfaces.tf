data "clouddk_network_interfaces" "example" {
  id = "${element(data.clouddk_servers.example.ids, 0)}"
}

output "data_clouddk_network_interfaces_example_addresses" {
  description = "The IP addresses assigned to the server's network interfaces"
  value       = "${data.clouddk_network_interfaces.example.addresses}"
}

output "data_clouddk_network_interfaces_example_default_firewall_rules" {
  description = "The default firewall rules for the server's network interfaces"
  value       = "${data.clouddk_network_interfaces.example.default_firewall_rules}"
}

output "data_clouddk_network_interfaces_example_firewall_rules_addresses" {
  description = "The commands for the firewall rules assigned to the server's network interfaces"
  value       = "${data.clouddk_network_interfaces.example.firewall_rules_addresses}"
}

output "data_clouddk_network_interfaces_example_firewall_rules_commands" {
  description = "The commands for the firewall rules assigned to the server's network interfaces"
  value       = "${data.clouddk_network_interfaces.example.firewall_rules_commands}"
}

output "data_clouddk_network_interfaces_example_firewall_rules_ids" {
  description = "The identifiers for the firewall rules assigned to the server's network interfaces"
  value       = "${data.clouddk_network_interfaces.example.firewall_rules_ids}"
}

output "data_clouddk_network_interfaces_example_firewall_rules_ports" {
  description = "The ports for the firewall rules assigned to the server's network interfaces"
  value       = "${data.clouddk_network_interfaces.example.firewall_rules_ports}"
}

output "data_clouddk_network_interfaces_example_firewall_rules_protocols" {
  description = "The protocols for the firewall rules assigned to the server's network interfaces"
  value       = "${data.clouddk_network_interfaces.example.firewall_rules_protocols}"
}

output "data_clouddk_network_interfaces_example_gateways" {
  description = "The gateways assigned to the server's network interfaces"
  value       = "${data.clouddk_network_interfaces.example.gateways}"
}

output "data_clouddk_network_interfaces_example_ids" {
  description = "The server's network interface identifiers"
  value       = "${data.clouddk_network_interfaces.example.ids}"
}

output "data_clouddk_network_interfaces_example_labels" {
  description = "The server's network interface labels"
  value       = "${data.clouddk_network_interfaces.example.labels}"
}

output "data_clouddk_network_interfaces_example_netmasks" {
  description = "The netmasks assigned to the server's network interfaces"
  value       = "${data.clouddk_network_interfaces.example.netmasks}"
}

output "data_clouddk_network_interfaces_example_networks" {
  description = "The networks assigned to the server's network interfaces"
  value       = "${data.clouddk_network_interfaces.example.networks}"
}

output "data_clouddk_network_interfaces_example_primary" {
  description = "Whether the network interface is the primary interface"
  value       = "${data.clouddk_network_interfaces.example.primary}"
}

output "data_clouddk_network_interfaces_example_rate_limits" {
  description = "The rate limits for the server's network interfaces"
  value       = "${data.clouddk_network_interfaces.example.rate_limits}"
}
