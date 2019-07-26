data "clouddk_servers" "example" {}

output "data_clouddk_servers_example_hostnames" {
  description = "The server hostnames"
  value       = "${data.clouddk_servers.example.hostnames}"
}

output "data_clouddk_servers_example_ids" {
  description = "The server identifiers"
  value       = "${data.clouddk_servers.example.ids}"
}

output "data_clouddk_servers_example_labels" {
  description = "The server labels"
  value       = "${data.clouddk_servers.example.labels}"
}

output "data_clouddk_servers_example_locations" {
  description = "The server location identifiers"
  value       = "${data.clouddk_servers.example.locations}"
}

output "data_clouddk_servers_example_packages" {
  description = "The server package identifiers"
  value       = "${data.clouddk_servers.example.packages}"
}

output "data_clouddk_servers_example_templates" {
  description = "The server template identifiers"
  value       = "${data.clouddk_servers.example.templates}"
}
#==============================================================================
data "clouddk_servers" "example_filter" {
  filter {
    hostname = "terraform"
  }
}

output "data_clouddk_servers_example_filter_hostnames" {
  description = "The server hostnames"
  value       = "${data.clouddk_servers.example_filter.hostnames}"
}

output "data_clouddk_servers_example_filter_ids" {
  description = "The server identifiers"
  value       = "${data.clouddk_servers.example_filter.ids}"
}

output "data_clouddk_servers_example_filter_labels" {
  description = "The server labels"
  value       = "${data.clouddk_servers.example_filter.labels}"
}

output "data_clouddk_servers_example_filter_locations" {
  description = "The server location identifiers"
  value       = "${data.clouddk_servers.example_filter.locations}"
}

output "data_clouddk_servers_example_filter_packages" {
  description = "The server package identifiers"
  value       = "${data.clouddk_servers.example_filter.packages}"
}

output "data_clouddk_servers_example_filter_templates" {
  description = "The server template identifiers"
  value       = "${data.clouddk_servers.example_filter.templates}"
}
