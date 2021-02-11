[![Build Status](https://api.travis-ci.com/danitso/terraform-provider-clouddk.svg?branch=master)](https://travis-ci.com/danitso/terraform-provider-clouddk)
[![Go Report Card](https://goreportcard.com/badge/github.com/danitso/terraform-provider-clouddk)](https://goreportcard.com/report/github.com/danitso/terraform-provider-clouddk)
[![GoDoc](https://godoc.org/github.com/danitso/terraform-provider-clouddk?status.svg)](http://godoc.org/github.com/danitso/terraform-provider-clouddk)

# Terraform Provider for Cloud.dk
A Terraform Provider which manages resources from [Cloud.dk](https://cloud.dk/).

## Requirements
- [Terraform](https://www.terraform.io/downloads.html) 0.13+
- [Go](https://golang.org/doc/install) 1.15+ (to build the provider plugin)
- [GoReleaser](https://goreleaser.com/install/) 0.155+ (to build the provider plugin)

## Table of Contents
- [Building the provider](#building-the-provider)
- [Using the provider](#using-the-provider)
- [Testing the provider](#testing-the-provider)

## Building the provider
- Clone the repository to `$GOPATH/src/github.com/danitso/terraform-provider-clouddk`:

    ```sh
    $ mkdir -p "${GOPATH}/src/github.com/danitso"
    $ cd "${GOPATH}/src/github.com/danitso"
    $ git clone git@github.com:danitso/terraform-provider-clouddk
    ```

- Enter the provider directory and build it:

    ```sh
    $ cd "${GOPATH}/src/github.com/danitso/terraform-provider-clouddk"
    $ make build
    ```

## Using the provider
You can find the latest release and its documentation in the [Terraform Registry](https://registry.terraform.io/providers/danitso/clouddk/latest).

## Testing the provider
In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

Tests are limited to regression tests, ensuring backwards compability.
