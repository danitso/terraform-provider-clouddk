# Terraform Provider for Cloud.dk
A Terraform Provider which manages resources from [Cloud.dk](https://cloud.dk/).

# Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.11+
- [Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)

# Building the Provider
Clone repository to: `$GOPATH/src/github.com/danitso/terraform-provider-clouddk`

```sh
$ mkdir -p $GOPATH/src/github.com/danitso; cd $GOPATH/src/github.com/danitso
$ git clone git@github.com:danitso/terraform-provider-clouddk
```

Enter the provider directory, initialize and build the provider

```sh
$ cd $GOPATH/src/github.com/danitso/terraform-provider-clouddk
$ make init
$ make build
```

# Using the Provider
If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-plugins) After placing it into your plugins directory,  run `terraform init` to initialize it.

## Configuration

### Arguments

* `endpoint` - (Optional) The API endpoint
* `key` - (Required) The API key

## Data Sources

### Disk (clouddk_disk)

#### Arguments

* `id` - (Required) This is the disk's identifier.
* `server_id` - (Required) This is the server's identifier.

#### Attributes

* `label` - This is the disk label.
* `primary` - Whether the disk is the primary disk.
* `size` - This is the disk size in gigabytes.

### Disks (clouddk_disks)

#### Arguments

* `id` - (Required) This is the server's identifier.

#### Attributes

* `ids` - This is the server's disk identifiers.
* `labels` - This is the server's disk labels.
* `primary` - Whether a disk is the primary disk.
* `sizes` - This is the server's disk sizes in gigabytes.

### Firewall Rule (clouddk_firewall_rule)

#### Arguments

* `id` - (Required) This is the firewall rule's identifier.
* `network_interface_id` - (Required) This is the network interface's identifier.
* `server_id` - (Required) This is the server's identifier.

#### Attributes

* `address` - This is the CIDR block for the firewall rule.
* `command` - This is the command for the firewall rule.
* `port` - This is the port for the firewall rule.
* `protocol` - This is the protocol for the firewall rule.

### Firewall Rules (clouddk_firewall_rules)

#### Arguments

* `id` - (Required) This is the network interface's identifier.
* `server_id` - (Required) This is the server's identifier.

#### Attributes

* `addresses` - This is the CIDR blocks for the firewall rules assigned to the network interface.
* `commands` - This is the commands for the firewall rules assigned to the network interface.
* `ids` - This is the identifiers for the firewall rules assigned to the network interface.
* `ports` - This is the ports for the firewall rules assigned to the network interface.
* `protocols` - This is the protocols for the firewall rules assigned to the network interface.

### IP Addresses (clouddk_ip_addresses)

#### Arguments

* `id` - (Required) This is the server's identifier.

#### Attributes

* `addresses` - This is the IP addresses assigned to the server's network interfaces.
* `gateways` - This is the gateways assigned to the server's network interfaces.
* `netmasks` - This is the netmasks assigned to the server's network interfaces.
* `network_interface_ids` - This is the network interface identifiers.
* `networks` - This is the networks assigned to the server's network interfaces.

### Locations (clouddk_locations)

#### Arguments

This data source has no arguments.

#### Attributes

* `ids` - This is the list of location identifiers.
* `names` - This is the list of location names.

### Network Interface (clouddk_network_interface)

#### Arguments

* `id` - (Required) This is the network interface's identifier.
* `server_id` - (Required) This is the server's identifier.

#### Attributes

* `addresses` - This is the IP addresses assigned to the network interface.
* `default_firewall_rule` - This is the default firewall rule for the network interface.
* `firewall_rules_addresses` - This is the CIDR blocks for the firewall rules assigned to the network interface.
* `firewall_rules_commands` - This is the commands for the firewall rules assigned to the network interface.
* `firewall_rules_ids` - This is the identifiers for the firewall rules assigned to the network interface.
* `firewall_rules_ports` - This is the ports for the firewall rules assigned to the network interface.
* `firewall_rules_protocols` - This is the protocols for the firewall rules assigned to the network interface.
* `gateways` - This is the gateways assigned to the network interface.
* `label` - This is the label for the network interface.
* `netmasks` - This is the netmasks assigned to the network interface.
* `networks` - This is the networks assigned to the network interface.
* `primary` - Whether a network interface is the primary interface.
* `rate_limit` - This is the rate limit for the network interface.

### Network Interfaces (clouddk_network_interfaces)

#### Arguments

* `id` - (Required) This is the server's identifier.

#### Attributes

* `addresses` - This is the IP addresses assigned to the server's network interfaces.
* `default_firewall_rules` - This is the default firewall rules for the server's network interfaces.
* `firewall_rules_addresses` - This is the CIDR blocks for the firewall rules assigned to the server's network interfaces.
* `firewall_rules_commands` - This is the commands for the firewall rules assigned to the server's network interfaces.
* `firewall_rules_ids` - This is the identifiers for the firewall rules assigned to the server's network interfaces.
* `firewall_rules_ports` - This is the ports for the firewall rules assigned to the server's network interfaces.
* `firewall_rules_protocols` - This is the protocols for the firewall rules assigned to the server's network interfaces.
* `gateways` - This is the gateways assigned to the server's network interfaces.
* `ids` - This is the server's network interface identifiers.
* `labels` - This is the server's network interface labels.
* `netmasks` - This is the netmasks assigned to the server's network interfaces.
* `networks` - This is the networks assigned to the server's network interfaces.
* `primary` - Whether a network interface is the primary interface.
* `rate_limits` - This is the rate limits for the server's network interfaces.

### Packages (clouddk_packages)

#### Arguments

This data source has no arguments.

#### Attributes

* `ids` - This is the list of package identifiers.
* `names` - This is the list of package names.

### Server (clouddk_server)

#### Arguments

* `id` - (Required) This is the server's identifier.

#### Attributes

* `booted` - Whether the server has been booted.
* `cpus` - This is the server's CPU count.
* `disk_ids` - This is the server's disk identifiers.
* `disk_labels` - This is the server's disk labels.
* `disk_primary` - Whether a disk is the primary disk.
* `disk_sizes` - This is the server's disk sizes in gigabytes.
* `hostname` - This is the server's hostname.
* `label` - This is the server's label.
* `location_id` - This is the location identifier.
* `location_name` - This is the location name.
* `memory` - This is the server's memory allocation in megabytes.
* `network_interface_addresses` - This is the IP addresses assigned to the server's network interfaces.
* `network_interface_default_firewall_rules` - This is the default firewall rules for the server's network interfaces.
* `network_interface_firewall_rules_addresses` - This is the CIDR blocks for the firewall rules assigned to the server's network interfaces.
* `network_interface_firewall_rules_commands` - This is the commands for the firewall rules assigned to the server's network interfaces.
* `network_interface_firewall_rules_ids` - This is the identifiers for the firewall rules assigned to the server's network interfaces.
* `network_interface_firewall_rules_ports` - This is the ports for the firewall rules assigned to the server's network interfaces.
* `network_interface_firewall_rules_protocols` - This is the protocols for the firewall rules assigned to the server's network interfaces.
* `network_interface_gateways` - This is the gateways assigned to the server's network interfaces.
* `network_interface_ids` - This is the server's network interface identifiers.
* `network_interface_labels` - This is the server's network interface labels.
* `network_interface_netmasks` - This is the netmasks assigned to the server's network interfaces.
* `network_interface_networks` - This is the networks assigned to the server's network interfaces.
* `network_interface_primary` - Whether a network interface is the primary interface.
* `network_interface_rate_limits` - This is the rate limits for the server's network interfaces.
* `package_id` - This is the package identifier.
* `package_name` - This is the package name.
* `template_id` - This is the template identifier.
* `template_name` - This is the template name.

### Servers (clouddk_servers)

#### Arguments

* `filter` - (Optional) This is the filter block.
    * `hostname` - (Optional) This is the hostname filter which performs a substring match on the hostname property.

#### Attributes

* `hostnames` - This is the list of server hostnames.
* `ids` - This is the list of server identifiers.
* `labels` - This is the list of server labels.
* `location_ids` - This is the list of server location identifiers.
* `location_names` - This is the list of server location names.
* `package_ids` - This is the list of server package identifiers.
* `package_names` - This is the list of server package names.
* `template_ids` - This is the list of server template identifiers.
* `template_names` - This is the list of server template names.

### Templates (clouddk_templates)

#### Arguments

* `filter` - (Optional) This is the filter block.
    * `name` - (Optional) This is the name filter which performs a substring match on the name property.

#### Attributes

* `ids` - This is the list of template identifiers.
* `names` - This is the list of template names.

## Resources

### Disk (clouddk_disk)

#### Arguments

* `label` - (Required) This is the disk label.
* `server_id` - (Required) This is the server's identifier.
* `size` - (Required) This is the disk size in gigabytes.

#### Attributes

* `id` - This is the disk's identifier.
* `primary` - Whether the disk is the primary disk.

### Firewall Rule (clouddk_firewall_rule)

#### Arguments

* `address` - (Required) This is the CIDR block for the firewall rule.
* `command` - (Required) This is the command for the firewall rule.
* `network_interface_id` - (Required) This is the network interface's identifier.
* `port` - (Required) This is the port for the firewall rule.
* `protocol` - (Required) This is the protocol for the firewall rule.
* `server_id` - (Required) This is the server's identifier.

#### Attributes

* `id` - This is the firewall rule's identifier.

### Server (clouddk_server)

#### Arguments

* `hostname` - (Required) This is the server's hostname.
* `label` - (Required) This is the server's label.
* `location_id` - (Required) This is the server's location.
* `package_id` - (Required) This is the server's package.
* `primary_network_interface_default_firewall_rule` - This is the default firewall rule for the server's primary network interface.
* `primary_network_interface_label` - This is the label for the server's primary network interface.
* `root_password` - (Required) This is the initial root password.
* `template_id` - (Required) This is the server's template.

#### Attributes

* `booted` - Whether the server has been booted.
* `cpus` - This is the server's CPU count.
* `disk_ids` - This is the server's disk identifiers.
* `disk_labels` - This is the server's disk labels.
* `disk_primary` - Whether a disk is the primary disk.
* `disk_sizes` - This is the server's disk sizes in gigabytes.
* `hostname` - This is the server's hostname.
* `id` - This is the server's identifier.
* `label` - This is the server's label.
* `location_id` - This is the location identifier.
* `location_name` - This is the location name.
* `memory` - This is the server's memory allocation in megabytes.
* `network_interface_addresses` - This is the IP addresses assigned to the server's network interfaces.
* `network_interface_default_firewall_rules` - This is the default firewall rules for the server's network interfaces.
* `network_interface_firewall_rules_addresses` - This is the CIDR blocks for the firewall rules assigned to the server's network interfaces.
* `network_interface_firewall_rules_commands` - This is the commands for the firewall rules assigned to the server's network interfaces.
* `network_interface_firewall_rules_ids` - This is the identifiers for the firewall rules assigned to the server's network interfaces.
* `network_interface_firewall_rules_ports` - This is the ports for the firewall rules assigned to the server's network interfaces.
* `network_interface_firewall_rules_protocols` - This is the protocols for the firewall rules assigned to the server's network interfaces.
* `network_interface_gateways` - This is the gateways assigned to the server's network interfaces.
* `network_interface_ids` - This is the server's network interface identifiers.
* `network_interface_labels` - This is the server's network interface labels.
* `network_interface_netmasks` - This is the netmasks assigned to the server's network interfaces.
* `network_interface_networks` - This is the networks assigned to the server's network interfaces.
* `network_interface_primary` - Whether a network interface is the primary interface.
* `network_interface_rate_limits` - This is the rate limits for the server's network interfaces.
* `package_id` - This is the package identifier.
* `package_name` - This is the package name.
* `template_id` - This is the template identifier.
* `template_name` - This is the template name.

# Developing the Provider
If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.12+ is *required*).
You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-clouddk
...
```

If you wish to contribute to the provider, the following requirements must be met,

* All tests must pass using `make test`
* The Go code must be formatted using Gofmt
* Dependencies are installed by `make init`

# Testing the Provider
In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

Tests are limited to regression tests, ensuring backwards compability.
