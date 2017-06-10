# Validate

This repository is to be used for some validation functions.

## Install

To install, use `go get`:

```bash
$ go get -u github.com/elpinal/validate/...
```

-u flag stands for "update".

## Examples

```go
package main

import (
	"fmt"
	"github.com/elpinal/validate/emails"
)

func main() {
	fmt.Printf("%v\n", emails.Validate("invalid..address@example.com"))
}
```

## Contribution

1. Fork ([https://github.com/elpinal/validate/fork](https://github.com/elpinal/validate/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[elpinal](https://github.com/elpinal)
