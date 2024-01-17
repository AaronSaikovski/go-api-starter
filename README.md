<div align="center">

# GoLang API Starter Template - v0.3.1

A simple Golang API project template that uses the Echo Web framework to save you time and energy.

[![Build Status](https://github.com/AaronSaikovski/go-api-starter/workflows/build/badge.svg)](https://github.com/AaronSaikovski/go-api-starter/actions)

</div>

A simple GoLang WebAPI boilerplate project to accelerate Golang projects. Originally based on [this](https://github.com/bmdavis419/the-better-backend) template. Full Credit to [Ben Davis](twitter.com/bmdavis419)

## This project contains:

- Echo Web framework - https://echo.labstack.com/
- Zerolog - https://github.com/rs/zerolog
- Swaggo - https://github.com/swaggo/swag
- CI (Github Actions)
- Unit tests (coming soon)
- Container support with [docker](Dockerfile) and [docker-compose](docker-compose.yml)
- API key validation - https://echo.labstack.com/docs/middleware/key-auth
- Compression - https://echo.labstack.com/docs/middleware/gzip
- Rate limiting - https://echo.labstack.com/docs/middleware/rate-limiter
- Auto Recovery - https://echo.labstack.com/docs/middleware/recover
- Logging - https://echo.labstack.com/docs/middleware/logger
- CORS Support - https://echo.labstack.com/docs/middleware/cors

## Install

Click the [Use this template](https://github.com/AaronSaikovski/go-api-starte/generate) button at the top of this project's GitHub page to get started.

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

View swagger docs at http://localhost:8080/swagger

### Extra packages

1. AIR server (supports hot reload) - `go install github.com/cosmtrek/air@latest`
2. Swagger - `go install github.com/swaggo/swag/cmd/swag@latest`

### Build and run

#### run `make help` for more assistance on the make file.

```
help          - Display help about make targets for this Makefile
localrelease  - Builds the project in preparation for (local)release
docs          - updates the swagger docs
debug         - Builds the project in preparation for debug
buildandrun   - builds and runs the program on the target platform
run           - runs main.go for testing
clean         - Remove the old builds and any debug information
unittest      - executes unit tests
dep           - fetches any external dependencies
vet           - Vet examines Go source code and reports suspicious constructs
staticcheck   - Runs static code analyzer staticcheck - currently broken
seccheck      - Code vulnerability check
lint          - format code and tidy modules
depupdate     - Update dependencies
hotload       - Uses the Air server for hot reloading support.
```

### Basic Ping Healthchecks

Uses a basic Ping/Pong approach for healthchecks.
Uses the endpoint:
```
http://127.0.0.1:8080/api/ping
```

Returns:

```
pong -
HTTP Status code - 200 OK:
```

If the API is running/healthy.

### Docker Support

This project has been fully tested with Docker and includes a `Dockerfile` and a `docker-compose.yml` file and uses a multi-stage docker file to reduce the docker image to approx. 30MB.
The docker build process also optimises the Go executable down to a bare minimum.
To build the docker image type:

```
docker build -t golangapi:1.0.0 .
```

and to run the comtainer type:

```
docker run -p 8080:8080 golangapi:1.0.0
```

or you can streamline the container build and run process by typing:

```
docker compose up
```

and to tear it down type:

```
docker compose down
```

### Issues

Please report any issues [here](https://github.com/AaronSaikovski/go-api-starter/issues).

## References

- [The Better backend](https://github.com/bmdavis419/the-better-backend)
- [Golang project Layout](https://github.com/golang-standards/project-layout)
- [The one-and-only, must-have, eternal Go project layout](https://appliedgo.com/blog/go-project-layout)
- [How To Upgrade Golang Dependencies](https://golang.cafe/blog/how-to-upgrade-golang-dependencies.html)

## Credits

[Ben Davis](https://twitter.com/bmdavis419) - [The better backend](https://github.com/bmdavis419/the-better-backend)
