data "clouddk_server_packages" "example" {}

output "data_clouddk_server_packages_example_ids" {
  description = "The package identifiers"
  value       = "${data.clouddk_server_packages.example.ids}"
}

output "data_clouddk_server_packages_example_names" {
  description = "The package names"
  value       = "${data.clouddk_server_packages.example.names}"
}
