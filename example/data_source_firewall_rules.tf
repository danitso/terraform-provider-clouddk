data "clouddk_firewall_rules" "example" {
  id = "${element(data.clouddk_network_interfaces.example.ids, 0)}"
  server_id = "${element(data.clouddk_servers.example.ids, 0)}"
}

output "data_clouddk_firewall_rules_example_addresses" {
  description = "The commands for the firewall rules assigned to the network interface"
  value       = "${data.clouddk_firewall_rules.example.addresses}"
}

output "data_clouddk_firewall_rules_example_commands" {
  description = "The commands for the firewall rules assigned to the network interface"
  value       = "${data.clouddk_firewall_rules.example.commands}"
}

output "data_clouddk_firewall_rules_example_ids" {
  description = "The identifiers for the firewall rules assigned to the network interface"
  value       = "${data.clouddk_firewall_rules.example.ids}"
}

output "data_clouddk_firewall_rules_example_ports" {
  description = "The ports for the firewall rules assigned to the network interface"
  value       = "${data.clouddk_firewall_rules.example.ports}"
}

output "data_clouddk_firewall_rules_example_protocols" {
  description = "The protocols for the firewall rules assigned to the network interface"
  value       = "${data.clouddk_firewall_rules.example.protocols}"
}
