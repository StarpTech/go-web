![big-gopher](big-gopher.png)

[![License MIT](https://img.shields.io/badge/License-MIT-blue.svg)](http://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/StarpTech/go-web.svg?branch=master)](https://travis-ci.org/StarpTech/go-web)
[![Go Report Card](https://goreportcard.com/badge/github.com/StarpTech/go-web)](https://goreportcard.com/report/github.com/StarpTech/go-web)

# Go-Web
Modern Web Application with Golang "Keep it simple, stupid"

# Stack
- HTTP Middleware [Echo](https://echo.labstack.com/)
- Language [Go](https://golang.org/) +1.8
- ORM library [gorm](https://github.com/jinzhu/gorm)
- Configuration [env](https://github.com/caarlos0/env)
- Load ENV variables from .env file [godotenv](https://github.com/joho/godotenv)
- Payload validation [validator](https://github.com/go-playground/validator)
- Templating [Go Templates](https://golang.org/pkg/text/template/)
- Frontend module bundler [rollupjs](https://rollupjs.org/)
- Stylesheet language [sass](http://sass-lang.com/)
- Redis Cache [Redis](https://github.com/go-redis/redis)
- Localization [gotext](https://github.com/leonelquinteros/gotext)

# Getting Started

```
$ go get ./...
$ docker-compose up
$ go run main.go
```

# Guidelines

- Javascript [standard](https://standardjs.com/)
- CSS [airbnb](https://github.com/airbnb/css)

# Tooling

- IDE plugin [vscode-go](https://github.com/Microsoft/vscode-go)
- Administration of postgresql database [pgAdmin](https://www.pgadmin.org/)
- REST client [Postman](https://chrome.google.com/webstore/detail/postman/fhbjgbiflinjbdggehcddcbncdddomop?hl=en)
- Go testing in the browser [go-convey](https://github.com/smartystreets/goconvey)
- Benchmarking [bombardier](http://github.com/codesenberg/bombardier)

# Benchmarking

```
$ bombardier -c 10 -n 10000 http://localhost:8080/feed
$ bombardier -c 10 -n 10000 http://localhost:8080/users/1
$ bombardier -c 10 -n 10000 http://localhost:8080/users/1/details
```

# TODO

- [X] Testsuite
- [X] Parameter Validation
- [X] Configuration
- [X] Load ENV variables from .env file
- [X] Add Graylog logger
- [X] Frontend build pipeline
- [X] Postgresql models
- [X] Gracefully shutdown
- [X] Consistent error-handling
- [X] Web Development with Custom elements
- [X] Templating
- [X] Correct linking with the app bundles
- [X] Metric endpoint (Prometheus)
- [X] Healthcheck endpoint
- [X] Distributed cache with Redis
- [X] Localization
- [ ] JWT Decoding
- [X] Benchmarks
- [X] Integration tests
- [X] Unit tests
- [X] Using interface for db, cache (better testable)
- [X] Setup Travis CI with static code analysis
- [ ] Swagger Documentation
- [ ] Code documentation `godoc`
- [ ] Rollup bundles with hashs

# Useful links

- http://www.alexedwards.net/blog/organising-database-access
