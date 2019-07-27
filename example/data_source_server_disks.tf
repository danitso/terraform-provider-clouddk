data "clouddk_server_disks" "example" {
  id = "${element(data.clouddk_servers.example.ids, 0)}"
}

output "data_clouddk_server_disks_example_ids" {
  description = "The server's disk identifiers"
  value       = "${data.clouddk_server_disks.example.ids}"
}

output "data_clouddk_server_disks_example_labels" {
  description = "The server's disk labels"
  value       = "${data.clouddk_server_disks.example.labels}"
}

output "data_clouddk_server_disks_example_primary" {
  description = "Whether a disk is the primary disk"
  value       = "${data.clouddk_server_disks.example.primary}"
}

output "data_clouddk_server_disks_example_sizes" {
  description = "The server's disk identifiers"
  value       = "${data.clouddk_server_disks.example.sizes}"
}
