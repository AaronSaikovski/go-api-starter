<div align="center">

# GoLang API Starter Template

A simple Golang API project template to save you time and energy.

[![Build Status](https://github.com/AaronSaikovski/go-api-starter/workflows/build/badge.svg)](https://github.com/AaronSaikovski/go-api-starter/actions)
[![Licence](https://img.shields.io/github/license/AaronSaikovski/go-api-starter)](LICENSE)

</div>
A simple GoLang WebAPI boilerplate project to accelerate Golang projects. Uses Fibre web framework and Zerolog for logging. This is based on the original template from Ben Davis - @bmdavis419.

## Install

Click the [Use this template](https://github.com/AaronSaikovski/go-api-starte/generate) button at the top of this project's GitHub page to get started.

View swagger docs at http://localhost:8080/swagger

## Usage

### Setup configuration

1. Configure the `go.mod` file and replace `module github.com/AaronSaikovski/go-api-starter` with your specific project url.
2. Configure the `Makefile` targets and parameters
3. Update the name in the `LICENSE` or swap it out entirely
4. Configure the `.github/workflows/build.yml` file
5. Update the `CHANGELOG.md` with your own info
6. Rename other files/folders as needed and configure their content
7. Delete this `README` and rename `README_project.md` to `README.md`
8. Run `go mod tidy` to ensure all the modules and packages are in place.
9. The build process is run from the `Makefile` and to test the project is working type: `make run` and check the console for output.

### Extra packages

1. AIR server (supports hot reload) - `go install github.com/cosmtrek/air@latest`
2. Swagger - `go install github.com/swaggo/swag/cmd/swag@latest`

### Build and run

#### run `make help` for more assistance on the make file.

1. `make build` - To make and build the program using the `Makefile`.
2. `make run` - To make and run the program using the `Makefile`.
3. `make test` - To make and run the unit tests using the `Makefile`.
4. `make clean` - To cleanup and delete all binaries using the `Makefile`.
5. `make lint` - To lint the code using `golangci-lint` via the `Makefile`.
6. `make dep` - To download all program dependencies using `Makefile`.
7. `make depupdate` - Upgrades all dependencies to the latest or minor patch release using `Makefile`.
8. `make docs` - Updates the swagger documents..

## References

- [The Better backend](https://github.com/bmdavis419/the-better-backend)
- [Golang project Layout](https://github.com/golang-standards/project-layout)
- [The one-and-only, must-have, eternal Go project layout](https://appliedgo.com/blog/go-project-layout)
- [How To Upgrade Golang Dependencies](https://golang.cafe/blog/how-to-upgrade-golang-dependencies.html)
