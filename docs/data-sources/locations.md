---
layout: page
title: clouddk_locations
permalink: /data-sources/locations
nav_order: 6
parent: Data Sources
---

# Data Source: clouddk_locations

Retrieves information about the available server locations.

## Example Usage

```
data "clouddk_locations" "example" {}
```

## Argument Reference

This data source has no arguments.

## Attribute Reference

* `ids` - This is the list of location identifiers.
* `names` - This is the list of location names.
