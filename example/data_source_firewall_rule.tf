data "clouddk_firewall_rule" "example" {
  id = "${element(data.clouddk_firewall_rules.example.ids, 0)}"
  network_interface_id = "${element(data.clouddk_network_interfaces.example.ids, 0)}"
  server_id = "${element(data.clouddk_servers.example.ids, 0)}"
}

output "data_clouddk_firewall_rule_example_address" {
  description = "The address for the firewall rule"
  value       = "${data.clouddk_firewall_rule.example.address}"
}

output "data_clouddk_firewall_rule_example_command" {
  description = "The command for the firewall rule"
  value       = "${data.clouddk_firewall_rule.example.command}"
}

output "data_clouddk_firewall_rule_example_port" {
  description = "The port for the firewall rule"
  value       = "${data.clouddk_firewall_rule.example.port}"
}

output "data_clouddk_firewall_rule_example_protocol" {
  description = "The protocols for the firewall rule"
  value       = "${data.clouddk_firewall_rule.example.protocol}"
}
