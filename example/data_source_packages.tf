data "clouddk_packages" "example" {}

output "data_clouddk_packages_example_ids" {
  description = "The package identifiers"
  value       = "${data.clouddk_packages.example.ids}"
}

output "data_clouddk_packages_example_names" {
  description = "The package names"
  value       = "${data.clouddk_packages.example.names}"
}
