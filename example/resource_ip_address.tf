resource "clouddk_ip_address" "example" {
  server_id = "${clouddk_server.example.id}"
}

output "resource_clouddk_ip_address_example_address" {
  description = "The IP address"
  value       = "${clouddk_ip_address.example.address}"
}

output "resource_clouddk_ip_address_example_gateway" {
  description = "The gateway address"
  value       = "${clouddk_ip_address.example.gateway}"
}

output "resource_clouddk_ip_address_example_id" {
  description = "The IP address identifier"
  value       = "${clouddk_ip_address.example.id}"
}

output "resource_clouddk_ip_address_example_netmask" {
  description = "The netmask"
  value       = "${clouddk_ip_address.example.netmask}"
}

output "resource_clouddk_ip_address_example_network" {
  description = "The network address"
  value       = "${clouddk_ip_address.example.network}"
}

output "resource_clouddk_ip_address_example_network_interface_id" {
  description = "The network interface identifier"
  value       = "${clouddk_ip_address.example.network_interface_id}"
}
