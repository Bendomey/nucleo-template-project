# Nucleo template: `project`
:mortar_board: Nucleo-based microservices project template.

## Features
- [ ] Nucleo v0.1.0 with full-detailed `nucleo.config.go` file.
- [ ] Common mono-repo project with a demo `greeter` service.
- [ ] Sample database `products` service (with file-based NeDB in development & MongoDB in production).
- [ ] Optional API Gateway service with detailed service settings.
- [ ] Beautiful static welcome page to test generated services & watch nodes and services.
- [ ] Optional Transporter & Cacher.
- [ ] Metrics & Tracing.
- [ ] Docker & Docker Compose & Kubernetes files.
- [ ] Unit tests with [Go Testing](https://pkg.go.dev/testing).
- [ ] Lint with [GolangCI-lint](https://github.com/golangci/golangci-lint).


## Install
> This is temporal.

Clone this template

```bash
$ git clone git@github.com:Bendomey/nucleo-template-project.git
```


## Make scripts
- `make run-dev`: Start development mode (load all services locally without transporter with hot-reload & REPL)
- `npm run`: Start production mode 
- `make lint`: Run GolangCI-lint
- `make test`: Run tests 

## License
TBD

## Contact
Copyright (c) 2023 Nucleo

[![@nucleo-go](https://img.shields.io/badge/github-nucleo-green.svg)](https://github.com/Bendomey/nucleo-go)