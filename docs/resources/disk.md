---
layout: page
title: clouddk_disk
permalink: /resources/disk
nav_order: 1
parent: Resources
---

# Resource: clouddk_disk

Manages a disk for a server.

## Example Usage

```
resource "clouddk_disk" "example" {
  server_id = clouddk_server.example.id
  label     = "Terraform Example"
  size      = 8
}
```

## Argument Reference

* `label` - (Required) This is the disk label.
* `server_id` - (Required) This is the server's identifier.
* `size` - (Required) This is the disk size in gigabytes.

## Attribute Reference

* `id` - This is the disk's identifier.
* `primary` - Whether the disk is the primary disk.
