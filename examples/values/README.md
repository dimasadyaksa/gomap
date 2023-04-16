## Values

The `Values` function is used to get the values of a map.
It takes a map as its argument and returns a slice of the values.

```go
package main

import (
	"fmt"

	"github.com/dimasadyaksa/gomap"
)

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	values := gomap.Values(m)

	for _, value := range values {
		fmt.Println(value)
	}
}

```

In this program, the `gomap.Values` function is called with the `m` map
as its argument, and the returned slice of values is assigned to the `values` variable.
The `values` variable is then iterated over and the values are printed to the console.

```bash
$ go run main.go
1
2
3
```

The order of the values in the slice is not guaranteed to be the same.