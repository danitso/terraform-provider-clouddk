data "clouddk_locations" "example" {}

output "data_clouddk_locations_example_ids" {
  description = "The location identifiers"
  value       = "${data.clouddk_locations.example.ids}"
}

output "data_clouddk_locations_example_names" {
  description = "The location names"
  value       = "${data.clouddk_locations.example.names}"
}
