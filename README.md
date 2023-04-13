# Gomap

[![codecov](https://codecov.io/github/dimasadyaksa/gomap/branch/develop/graph/badge.svg?token=TNDBN2DH6G)](https://codecov.io/github/dimasadyaksa/gomap)

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

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.


## License
[MIT](https://choosealicense.com/licenses/mit/)

