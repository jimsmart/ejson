# ejson

[![BSD-style](https://img.shields.io/badge/license-BSD--style-blue.svg?style=flat)](LICENSE)
[![Build Status](https://github.com/jimsmart/ejson/actions/workflows/main.yml/badge.svg?branch=master)](https://github.com/jimsmart/ejson/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/jimsmart/ejson/branch/master/graph/badge.svg)](https://codecov.io/gh/jimsmart/ejson)
[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/jimsmart/ejson)

ejson is a [Go](https://golang.org) package that implements encoding and decoding of Meteor's EJSON (Extended JSON).

The mapping between EJSON and Go values is identical to package encoding/json (which this package is based upon), but with built-in EJSON encoding and decoding for `[]byte` and `time.Time`.

All EJSON serializations are also valid JSON. For example, an object with a date (`time.Time`) and some binary data (`[]byte`) would be serialized in EJSON as:

```json
{
    "d": {"$date": 1358205756553},
    "b": {"$binary": "c3VyZS4="}
}
```

Where `$date` is Unix time in milliseconds, and `$binary` is base64 encoded data.

This package should be used as a drop-in replacement for package encoding/json whenever EJSON encoding is required.

Package ejson is derived from a subtree-branch of Go's encoding/json package, currently tracking Go 1.7.4.

## Installation

```bash
go get github.com/jimsmart/ejson
```

```go
import "github.com/jimsmart/ejson"
```

## Documentation

GoDocs [https://godoc.org/github.com/jimsmart/ejson](https://godoc.org/github.com/jimsmart/ejson)

## Testing

To run the tests execute `go test` inside the project folder.

## License

Package ejson consists of minor modifications to Go standard library's encoding/json package. These modifications are copyright 2016 Jim Smart and released under a BSD-style license.

Package encoding/json is copyright 2009 The Go Authors and released under a [BSD-style license](LICENSE).

## Contributing

Bug reports are helpful and pull requests are welcome.

## History

- v0.0.1 (2021-04-19) Use Go modules. Switch from Travis CI to GitHub Actions.
