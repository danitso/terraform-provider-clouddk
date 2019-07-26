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
  description = "The server's disk identifiers"
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

output "data_clouddk_server_example_memory" {
  description = "The server's memory allocation in gigabytes"
  value       = "${data.clouddk_server.example.memory}"
}
