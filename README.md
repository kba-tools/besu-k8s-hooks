# Besu-K8s-Hooks

Custom hooks made for Besu to use in Kubernetes clusters.

[![Go Version](https://img.shields.io/github/go-mod/go-version/kba-tools/besu-k8s-hooks)](./go.mod)
[![Go Reference](https://pkg.go.dev/badge/github.com/kba-tools/besu-k8s-hooks.svg)](https://pkg.go.dev/github.com/kba-tools/besu-k8s-hooks)
[![Go Report Card](https://goreportcard.com/badge/github.com/kba-tools/besu-k8s-hooks)](https://goreportcard.com/report/github.com/kba-tools/besu-k8s-hooks)
[![Release](https://img.shields.io/github/v/release/kba-tools/besu-k8s-hooks)](https://github.com/kba-tools/besu-k8s-hooks/releases)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE.md)

## Hooks

### Besu-Config-Generator

Install:

```sh
curl -sSfL https://raw.githubusercontent.com/kba-tools/besu-k8s-hooks/main/install.sh | sh -s
```

Help:

```sh
./bin/besu-config-generator -h
```

Example run:

```sh
./bin/besu-config-generator  --xemptyBlockPeriod true  --emptyBlockPeriod 300   --validators 4 --chainID 1337 --blockperiod 15 --epochLength 30000 --requestTimeout 300 --gasLimit 0xfff --coinbase 0x0000000000000000000000000000000000000000  --accountPassword pwd --output ./bin/config
```
