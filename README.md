# go-http-client

[![Build](https://github.com/NdoleStudio/go-http-client/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/go-http-client/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/go-http-client/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/go-http-client)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/go-http-client/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/go-http-client/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/go-http-client)](https://goreportcard.com/report/github.com/NdoleStudio/go-http-client)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/go-http-client)](https://github.com/NdoleStudio/go-http-client/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/go-http-client?color=brightgreen)](https://github.com/NdoleStudio/go-http-client/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/go-http-client)](https://pkg.go.dev/github.com/NdoleStudio/go-http-client)


This package provides a generic `go` client template for an HTTP API

## Installation

`go-http-client` is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NdoleStudio/go-http-client
```

Alternatively the same can be achieved if you use `import` in a package:

```go
import "github.com/NdoleStudio/go-http-client"
```


## Implemented

- [Status Codes](#status-codes)
    - `GET /200`: OK

## Usage

### Initializing the Client

An instance of the client can be created using `New()`.

```go
package main

import (
	"github.com/NdoleStudio/go-http-client"
)

func main()  {
	statusClient := client.New(client.WithDelay(200))
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
status, response, err := statusClient.Status.Ok(context.Background())
if err != nil {
    //handle error
}
```

### Status Codes

#### `GET /200`: OK

```go
status, response, err := statusClient.Status.Ok(context.Background())

if err != nil {
    log.Fatal(err)
}

log.Println(status.Description) // OK
```

## Testing

You can run the unit tests for this client from the root directory using the command below:

```bash
go test -v
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
