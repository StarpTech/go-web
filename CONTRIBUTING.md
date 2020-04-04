## Setup your machine

`go-web` is written in [Go](https://golang.org/).

Prerequisites:

- `make`
- [Go 1.14+](https://golang.org/doc/install)
- [Docker](https://www.docker.com/)
- `gpg` (probably already installed on your system)

Clone `goweb` anywhere:

```sh
$ git clone git@github.com:StarpTech/go-web.git
```

Install the build and lint dependencies:

```sh
$ make setup
```

A good way of making sure everything is all right is running the test suite:

```sh
$ make test
```

## Test your change

You can create a branch for your changes and try to build from the source as you go:

```sh
$ make
```

When you are satisfied with the changes, we suggest you run:

```sh
$ make ci
```

Which runs all the linters and tests.

## Create a commit

Commit messages should be well formatted, and to make that "standardized", we
are using Conventional Commits.

You can follow the documentation on
[their website](https://www.conventionalcommits.org).

## Submit a pull request

Push your branch to your `go-web` fork and open a pull request against the
master branch.

## Deployment

Tag a new release and push it to origin. This will trigger the Github CI to deploy the commit.
```
git tag v1.0.0
git push origin v1.0.0
```
