provider "clouddk" {
  key = "${var.key}"
}

variable "key" {
  description = "The API key"
}
