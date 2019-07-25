data "clouddk_locations" "example" {}

output "data_clouddk_locations_example_result" {
  description = "The locations"
  value       = "${data.clouddk_locations.example.result}"
}
