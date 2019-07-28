resource "clouddk_disk" "example" {
  label = "Terraform Example"
  size = 16

  server_id = "${clouddk_server.example.id}"
}

output "resource_clouddk_disk_example_id" {
  description = "The disk identifier"
  value       = "${clouddk_disk.example.id}"
}

output "resource_clouddk_disk_example_label" {
  description = "The disk label"
  value       = "${clouddk_disk.example.label}"
}

output "resource_clouddk_disk_example_primary" {
  description = "Whether the disk is the primary disk"
  value       = "${clouddk_disk.example.primary}"
}

output "resource_clouddk_disk_example_size" {
  description = "The disk size in gigabytes"
  value       = "${data.clouddk_disk.example.size}"
}
