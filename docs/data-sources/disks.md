---
layout: page
title: clouddk_disks
permalink: /data-sources/disks
nav_order: 2
parent: Data Sources
---

# Data Source: clouddk_disks

Retrieves information about the disks for a server.

## Example Usage

```
data "clouddk_disks" "example" {
  id = clouddk_server.example.id
}
```

## Argument Reference

* `id` - (Required) This is the server's identifier.

## Attribute Reference

* `ids` - This is the server's disk identifiers.
* `labels` - This is the server's disk labels.
* `primary` - Whether a disk is the primary disk.
* `sizes` - This is the server's disk sizes in gigabytes.
