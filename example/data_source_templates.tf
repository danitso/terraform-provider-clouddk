data "clouddk_templates" "example" {}

output "data_clouddk_templates_example_ids" {
  description = "The template identifiers"
  value       = "${data.clouddk_templates.example.ids}"
}

output "data_clouddk_templates_example_names" {
  description = "The template names"
  value       = "${data.clouddk_templates.example.names}"
}

data "clouddk_templates" "example_filter" {
  filter {
    name = "ubuntu"
  }
}

output "data_clouddk_templates_example_filter_ids" {
  description = "The template identifiers"
  value       = "${data.clouddk_templates.example_filter.ids}"
}

output "data_clouddk_templates_example_filter_names" {
  description = "The template names"
  value       = "${data.clouddk_templates.example_filter.names}"
}
