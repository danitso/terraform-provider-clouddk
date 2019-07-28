data "clouddk_firewall_rule" "example" {
  id                   = "${clouddk_firewall_rule.example.id}"
  network_interface_id = "${clouddk_firewall_rule.example.network_interface_id}"
  server_id            = "${clouddk_server.example.id}"
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
