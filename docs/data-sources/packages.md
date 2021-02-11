---
layout: page
title: clouddk_packages
permalink: /data-sources/packages
nav_order: 9
parent: Data Sources
---

# Data Source: clouddk_packages

Retrieves information about the available server packages.

## Example Usage

```
data "clouddk_packages" "example" {}
```

## Argument Reference

This data source has no arguments.

## Attribute Reference

* `ids` - This is the list of package identifiers.
* `names` - This is the list of package names.
