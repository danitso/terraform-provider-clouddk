---
layout: home
title: Introduction
nav_order: 1
---

# Cloud.dk Provider

This provider for [Terraform](https://www.terraform.io/) is used for interacting with resources supported by [Cloud.dk](https://cloud.dk). The provider needs to be configured with the proper endpoint and key before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```
provider "clouddk" {
  key = "your-api-key-here"
}
```

## Argument Reference

* `endpoint` - (Optional) The API endpoint (defaults to `https://api.cloud.dk/v1`)
* `key` - (Required) The API key
