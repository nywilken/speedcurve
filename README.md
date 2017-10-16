# Speedcurve

## Summary
The speedcurve package is a Go based wrapper around Speedcurve's API.
For Speedcurve's source API documentation see https://api.speedcurve.com

## Installation

```bash
go get github.com/nywilken/speedcurve
```
## Usage

```go

package main

import (
	"fmt"
	"os"

	"github.com/nywilken/speedcurve"
)

func main() {
	token, _ := os.LookupEnv("SPD_API_TOKEN")
	sc := speedcurve.NewClient(token, "")

	d, _ := sc.GetLatestDeploy()
	fmt.Printf("%v", d)

	// Get TestID for all competed tests
	for _, t := range d.TestsCompleted {
		res, _ := sc.GetTest(t.Test)
		fmt.Printf("%v", res)
	}
}

```

## BUGS/PROBLEMS/CONTRIBUTING
If you do find something that doesn't work as expected, please file an issue on
Github: <https://github.com/nywilken/speedcurve/issues>

