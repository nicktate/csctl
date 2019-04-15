# csctl - The Official Containership CLI

[![CircleCI](https://circleci.com/gh/containership/csctl.svg?style=svg)](https://circleci.com/gh/containership/csctl)
[![Go Report Card](https://goreportcard.com/badge/github.com/containership/csctl)](https://goreportcard.com/report/github.com/containership/csctl)
[![codecov](https://codecov.io/gh/containership/csctl/branch/master/graph/badge.svg)](https://codecov.io/gh/containership/csctl)

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](code-of-conduct.md)

```
csctl is a command line interface for Containership.

Find more information at: https://github.com/containership/csctl

Usage:
  csctl [command]

Available Commands:
  create      Create a resource
  delete      Delete a resource
  export      Export a resource
  get         Get a resource
  help        Help about any command
  scale       Scale a node pool
  ssh         SSH into nodes of a cluster
  upgrade     Upgrade a cluster node pool or resource
  version     Output the current version of csctl

Flags:
      --config string   config file (default is ~/.containership/csctl.yaml)
      --debug           enable/disable debug mode (trace all HTTP requests)
  -h, --help            help for csctl
      --token string    Containership token to authenticate with

Use "csctl [command] --help" for more information about a command.
```

This repository also contains the go client for interacting with Containership Cloud.
For more info, please refer to [the client documentation](cloud/README.md).

**Warning**: This project is currently under active development and is subject to breaking changes without notice.

## Installing

To install the `csctl` binary, simply run:

```
go get -u github.com/containership/csctl
```

Alternatively, you can clone this repository and install via `make`:

```
make install
```

### Docker

Docker images are also provided for users that don't want to install the binary.
For example, to run the `latest` tag:

```
docker run containership/csctl <args>
```

Here's a convenient shell alias for running the `latest` Docker image with the config file at the default location mounted in:

```
alias csctl="docker run \
            --mount type=bind,source=$HOME/.containership/csctl.yaml,target=/app/csctl.yaml \
            containership/csctl --config /app/csctl.yaml"
```

Now it can be invoked as `csctl <args>` as usual.

## Usage

Please use `csctl -h` for now to discover usage.

More documentation will be added as the project matures.

### Configuration

`csctl` defaults to a config file located at `~/.containership/csctl.yaml`.
You may also choose to manually specify a config file using `--config`.

#### Authentication

A token is required to authenticate with Containership Cloud for almost every command.

It's recommended to generate a new [Personal Access Token](https://docs.containership.io/developer-resources/personal-access-tokens) through the UI for this.
`csctl` can then be configured to use this token by adding it under the `token` key in the config file.
The token can also be specified on the command line via `--token` for any command that requires it; however, setting the `token` key of the config file is the preferred approach to limit command line complexity.

#### Contexts

When interacting with Containership resources, it's often required to specify an organization, cluster, etc.
Instead of supplying these options via command lines flags every time, it's possible to specify them in the config file, creating a basic context to operate within.
Command line flags will override config file keys of the same name.

For example, one way to list all clusters in an organization is to supply the `--organization` flag on the command line as follows:

```
csctl get clusters --organization a4de3b04-60eb-45c2-a5d5-0d145fa1de58
```

However, chances are that most of the work a user does is within the context of a single organization.
In this case, a better way to specify the organization context is through the config file.
The most basic config file that can accomplish the same as above but without any command line flags is the following:

```
token: <containership_cloud_token>
organization: a4de3b04-60eb-45c2-a5d5-0d145fa1de58
```

The same context concept extends to clusters and node pools.

## Contributing

Thank you for your interest in this project and for your interest in contributing!
Feel free to open issues for feature requests, bugs, or even just questions - we love feedback and want to hear from you.

PRs are also always welcome!
However, if the feature you're considering adding is fairly large in scope, please consider opening an issue for discussion first.
