data "clouddk_templates" "example" {}

output "data_clouddk_templates_example_result" {
  description = "The templates"
  value       = "${data.clouddk_templates.example.result}"
}
