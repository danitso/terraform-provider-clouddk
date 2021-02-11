---
layout: page
title: clouddk_disk
permalink: /data-sources/disk
nav_order: 1
parent: Data Sources
---

# Data Source: clouddk_disk

Retrieves information about a disk for a server.

## Example Usage

```
data "clouddk_disk" "example" {
  id        = element(flatten(clouddk_server.example.disk_ids), 0)
  server_id = clouddk_server.example.id
}
```

## Argument Reference

* `id` - (Required) This is the disk's identifier.
* `server_id` - (Required) This is the server's identifier.

## Attribute Reference

* `label` - This is the disk label.
* `primary` - Whether the disk is the primary disk.
* `size` - This is the disk size in gigabytes.
