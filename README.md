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

### Locations (clouddk_locations)

#### Arguments

This data source has no arguments.

#### Attributes

* `ids` - This is the list of project ids.
* `names` - This is the list of project names.

### Templates (clouddk_templates)

#### Arguments

* `filter` - (Optional) This is the filter block.
    * `name` - (Optional) This is the name filter which removes matches which does not contain this string.

#### Attributes

* `ids` - This is the list of project ids.
* `names` - This is the list of project names.

## Resources

### ... (clouddk_...)

#### Arguments

* `undefined` - (Required) Work in progress.

#### Attributes

* `undefined` - Work in progress.

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
