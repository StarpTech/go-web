![big-gopher](big-gopher.png)

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

# Getting Started

```
$ go get ./...
$ docker-compose up
$ go run main.go
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
- [ ] Correct linking with the app bundles
- [X] Metric endpoint (Prometheus)
- [X] Healthcheck endpoint (Prometheus)
- [ ] JWT Decoding
- [ ] Benchmarks
- [ ] Unit tests
- [X] Setup Travis CI with static code analysis
- [ ] Swagger Documentation
- [ ] Code documentation `godoc`
