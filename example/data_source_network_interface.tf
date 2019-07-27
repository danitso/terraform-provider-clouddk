data "clouddk_network_interface" "example" {
  id = "${element(data.clouddk_servers.example.ids, 0)}"
  network_interface_id = "${element(data.clouddk_network_interfaces.example.ids, 0)}"
}

output "data_clouddk_network_interface_example_addresses" {
  description = "The IP addresses assigned to the network interface"
  value       = "${data.clouddk_network_interface.example.addresses}"
}

output "data_clouddk_network_interface_example_default_firewall_rule" {
  description = "The default firewall rule for the network interface"
  value       = "${data.clouddk_network_interface.example.default_firewall_rule}"
}

output "data_clouddk_network_interface_example_firewall_rules_addresses" {
  description = "The commands for the firewall rules assigned to the network interface"
  value       = "${data.clouddk_network_interface.example.firewall_rules_addresses}"
}

output "data_clouddk_network_interface_example_firewall_rules_commands" {
  description = "The commands for the firewall rules assigned to the network interface"
  value       = "${data.clouddk_network_interface.example.firewall_rules_commands}"
}

output "data_clouddk_network_interface_example_firewall_rules_ids" {
  description = "The identifiers for the firewall rules assigned to the network interface"
  value       = "${data.clouddk_network_interface.example.firewall_rules_ids}"
}

output "data_clouddk_network_interface_example_firewall_rules_ports" {
  description = "The ports for the firewall rules assigned to the network interface"
  value       = "${data.clouddk_network_interface.example.firewall_rules_ports}"
}

output "data_clouddk_network_interface_example_firewall_rules_protocols" {
  description = "The protocols for the firewall rules assigned to the network interface"
  value       = "${data.clouddk_network_interface.example.firewall_rules_protocols}"
}

output "data_clouddk_network_interface_example_gateways" {
  description = "The gateways assigned to the network interface"
  value       = "${data.clouddk_network_interface.example.gateways}"
}

output "data_clouddk_network_interface_example_label" {
  description = "The network interface's label"
  value       = "${data.clouddk_network_interface.example.label}"
}

output "data_clouddk_network_interface_example_netmasks" {
  description = "The netmasks assigned to the network interface"
  value       = "${data.clouddk_network_interface.example.netmasks}"
}

output "data_clouddk_network_interface_example_networks" {
  description = "The networks assigned to the network interface"
  value       = "${data.clouddk_network_interface.example.networks}"
}

output "data_clouddk_network_interface_example_primary" {
  description = "Whether the network interface is the primary interface"
  value       = "${data.clouddk_network_interface.example.primary}"
}

output "data_clouddk_network_interface_example_rate_limit" {
  description = "The rate limit for the network interface"
  value       = "${data.clouddk_network_interface.example.rate_limit}"
}
