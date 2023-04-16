## Reduce

The `Reduce` function is used to reduce a map to a single value based on a reducer function. 
It takes a map, an initial value, and a reducer function as its arguments. 
The reducer function takes the accumulator, the key, and the value of the map as its arguments.

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

	sumReducer := func(acc int, key string, value int) int {
		return acc + value
	}

	sum := gomap.Reduce(m, 0, sumReducer)

	fmt.Println(sum)
}
```

In this program, the `gomap.Reduce` function is called with the `m` map as its argument, 
an initial value of 0, and a `sumReducer` function that adds the value to the accumulator. 
The returned value is assigned to the `sum` variable and printed to the console.

```bash
$ go run main.go
6
```