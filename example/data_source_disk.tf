data "clouddk_disk" "example" {
  id = "${element(data.clouddk_disks.example.ids, 0)}"
  server_id = "${element(data.clouddk_servers.example.ids, 0)}"
}

output "data_clouddk_disk_example_label" {
  description = "The disk label"
  value       = "${data.clouddk_disk.example.label}"
}

output "data_clouddk_disk_example_primary" {
  description = "Whether the disk is the primary disk"
  value       = "${data.clouddk_disk.example.primary}"
}

output "data_clouddk_disk_example_size" {
  description = "The disk size in gigabytes"
  value       = "${data.clouddk_disk.example.size}"
}
