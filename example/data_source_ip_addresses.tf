data "clouddk_ip_addresses" "example" {
  id = "${clouddk_server.example.id}"
}

output "data_clouddk_ip_addresses_example_addresses" {
  description = "The IP addresses assigned to the server's network interfaces"
  value       = "${data.clouddk_ip_addresses.example.addresses}"
}

output "data_clouddk_ip_addresses_example_gateways" {
  description = "The gateways assigned to the server's network interfaces"
  value       = "${data.clouddk_ip_addresses.example.gateways}"
}

output "data_clouddk_ip_addresses_example_netmasks" {
  description = "The netmasks assigned to the server's network interfaces"
  value       = "${data.clouddk_ip_addresses.example.netmasks}"
}

output "data_clouddk_ip_addresses_example_network_interface_ids" {
  description = "The network interface identifiers"
  value       = "${data.clouddk_ip_addresses.example.network_interface_ids}"
}

output "data_clouddk_ip_addresses_example_networks" {
  description = "The networks assigned to the server's network interfaces"
  value       = "${data.clouddk_ip_addresses.example.networks}"
}
