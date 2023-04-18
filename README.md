# Gomap

[![codecov](https://codecov.io/gh/dimasadyaksa/gomap/branch/main/graph/badge.svg?token=TNDBN2DH6G)](https://codecov.io/gh/dimasadyaksa/gomap)
[![Go Report Card](https://goreportcard.com/badge/github.com/dimasadyaksa/gomap)](https://goreportcard.com/report/github.com/dimasadyaksa/gomap)
[![Go Reference](https://pkg.go.dev/badge/github.com/dimasadyaksa/gomap.svg)](https://pkg.go.dev/github.com/dimasadyaksa/gomap)
![GitHub](https://img.shields.io/github/license/dimasadyaksa/gomap)

Gomap is a package that contains several functions to make it easier to work with maps in Go.

## Installation

```bash
go get github.com/dimasadyaksa/gomap
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/dimasadyaksa/gomap"
)

func main(){
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    
    keys := gomap.Keys(m)
    fmt.Println(keys) // [a b c]
    
    values := gomap.Values(m)
    fmt.Println(values) // [1 2 3]
}
```
Please refer to the [examples](https://github.com/dimasadyaksa/gomap/tree/main/examples) directory for more usage examples.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.


## License
[MIT](https://choosealicense.com/licenses/mit/)

