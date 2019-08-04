provider "clouddk" {
  key = "${var.key}"
}

provider "random" {
  version = "~> 2.1"
}

variable "key" {
  description = "The API key"
}
