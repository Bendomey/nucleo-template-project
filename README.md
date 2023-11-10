# Nucleo template: `project`
:mortar_board: Nucleo-based microservices project template.

## Features
- [x] Nucleo v0.1.1 with full-detailed `configs/nucleo.config.go` file.
- [x] Common mono-repo project with a demo `greeter` service.
- [ ] Sample database `products` service (with file-based NeDB in development & MongoDB in production).
- [ ] Optional API Gateway service with detailed service settings.
- [ ] Beautiful static welcome page to test generated services & watch nodes and services.
- [ ] Optional Transporter & Cacher.
- [ ] Metrics & Tracing.
- [ ] Docker & Docker Compose & Kubernetes files.
- [x] Unit tests with [Go Testing](https://pkg.go.dev/testing).
- [x] Lint with [GolangCI-lint](https://golangci-lint.run/).


## Install
> This is temporal.

Clone this template

```bash
$ git clone git@github.com:Bendomey/nucleo-template-project.git
```


## Make scripts
- `make run-dev`: Start app with hot-reload
- `make run`: Start app without hot-reload
- `make lint`: Run GolangCI-lint
- `make utest`: Run unit tests 
- `make itest`: Run integration tests

## License
nucleo-template-project is available under the [Apache License](https://www.tldrlegal.com/license/apache-license-2-0-apache-2-0)

## Contact
Copyright (c) 2023 Nucleo

[![@nucleo-go](https://img.shields.io/badge/github-nucleo-green.svg)](https://github.com/Bendomey/nucleo-go)