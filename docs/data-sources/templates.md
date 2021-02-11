---
layout: page
title: clouddk_templates
permalink: /data-sources/templates
nav_order: 12
parent: Data Sources
---

# Data Source: clouddk_templates

Retrieves information about the available templates.

## Example Usage

```
data "clouddk_templates" "example" {}
```

## Argument Reference

* `filter` - (Optional) This is the filter block.
    * `name` - (Optional) This is the name filter which performs a substring match on the name property.

## Attribute Reference

* `ids` - This is the list of template identifiers.
* `names` - This is the list of template names.
