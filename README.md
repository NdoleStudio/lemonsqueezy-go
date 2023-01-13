# lemonsqueezy-go

[![Build](https://github.com/NdoleStudio/lemonsqueezy-go/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/lemonsqueezy-go/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/lemonsqueezy-go/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/lemonsqueezy-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/lemonsqueezy-go/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/lemonsqueezy-go/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/lemonsqueezy-go)](https://goreportcard.com/report/github.com/NdoleStudio/lemonsqueezy-go)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/lemonsqueezy-go)](https://github.com/NdoleStudio/lemonsqueezy-go/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/lemonsqueezy-go?color=brightgreen)](https://github.com/NdoleStudio/lemonsqueezy-go/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/lemonsqueezy-go)](https://pkg.go.dev/github.com/NdoleStudio/lemonsqueezy-go)


This package provides a go API client for the lemonsqueezy API

## Installation

`lemonsqueezy-go` is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NdoleStudio/lemonsqueezy-go
```

Alternatively the same can be achieved if you use `import` in a package:

```go
import "github.com/NdoleStudio/lemonsqueezy-go"
```

## Implemented

- **Users**
  - `GET /v1/users/me`: Retrieves the currently authenticated user.
- **Stores**
  - `GET /v1/stores/:id`: Retrieve a store
  - `GET /v1/stores`: List all stores
- **Subscriptions**
  - `PATCH /v1/subscriptions/:id`: Update a subscription
  - `GET /v1/subscriptions/:id`: Retrieve a subscription
  - `GET /v1/subscriptions`: List all subscriptions
  - `DELETE /v1/subscriptions/{id}`: Cancel an active subscription
- **Webhooks**
  - `Verify`: Verify that webhook requests are coming from Lemon Squeezy

## Usage

### Initializing the Client

An instance of the client can be created using `New()`.

```go
package main

import (
    "github.com/NdoleStudio/lemonsqueezy-go"
)

func main() {
    client := lemonsqueezy.New(lemonsqueezy.WithAPIKey(""))
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
subscription, response, err := client.Subscriptions.Get(context.Background(), "1")
if err != nil {
    //handle error
}
```

## Testing

You can run the unit tests for this client from the root directory using the command below:

```bash
go test -v
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
