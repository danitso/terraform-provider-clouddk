---
layout: page
title: clouddk_servers
permalink: /data-sources/servers
nav_order: 11
parent: Data Sources
---

# Data Source: clouddk_servers

Retrieves information about the available servers.

## Example Usage

```
data "clouddk_servers" "example" {}
```

## Argument Reference

* `filter` - (Optional) This is the filter block.
    * `hostname` - (Optional) This is the hostname filter which performs a substring match on the hostname property.

## Attribute Reference

* `hostnames` - This is the list of server hostnames.
* `ids` - This is the list of server identifiers.
* `labels` - This is the list of server labels.
* `location_ids` - This is the list of server location identifiers.
* `location_names` - This is the list of server location names.
* `package_ids` - This is the list of server package identifiers.
* `package_names` - This is the list of server package names.
* `template_ids` - This is the list of server template identifiers.
* `template_names` - This is the list of server template names.
