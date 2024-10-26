# shup

[![Build Status](https://github.com/ai-shflow/shup/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/ai-shflow/shup/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/ai-shflow/shup/branch/main/graph/badge.svg?token=El8oiyaIsD)](https://codecov.io/gh/ai-shflow/shup)
[![Go Report Card](https://goreportcard.com/badge/github.com/ai-shflow/shup)](https://goreportcard.com/report/github.com/ai-shflow/shup)
[![License](https://img.shields.io/github/license/ai-shflow/shup.svg)](https://github.com/ai-shflow/shup/blob/main/LICENSE)
[![Tag](https://img.shields.io/github/tag/ai-shflow/shup.svg)](https://github.com/ai-shflow/shup/tags)



## Introduction

*shup* is the installer of [shai](https://github.com/ai-shflow/shai) written in Go.



## Prerequisites

- Go >= 1.22.0



## Build

```bash
version=latest make build
```



## Usage

```
shai installer

Usage:
  shup [flags]
  shup [command]

Available Commands:
  check       Check for updates to toolchains and shup
  help        Help about any command
  show        Show the installed toolchains
  update      Update toolchains and shup

Flags:
      --config string   config file (default "$HOME/.shai/shup.yml")
  -h, --help            help for shup
  -v, --version         version for shup

Use "shup [command] --help" for more information about a command.
```



## Settings

*shup* parameters can be set in the directory [config](https://github.com/ai-shflow/shup/blob/main/config).

An example of configuration in [config.yml](https://github.com/ai-shflow/shup/blob/main/config/config.yml):

```yaml
apiVersion: v1
kind: shup
metadata:
  name: shup
spec:
  artifact:
    url: http://127.0.0.1:8080
    user: user
    pass: pass
```



## License

Project License can be found [here](LICENSE).



## Reference

- [bubbletea](https://github.com/charmbracelet/bubbletea/)
