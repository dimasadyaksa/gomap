## Some

The `Some` function is used to check if some values in a map satisfy a predicate function. It takes a map and a predicate function as its arguments.

```go
package main

import (
	"fmt"
	"github.com/dimasadyaksa/gomap"
)

func main() {
	m := map[string]int{
		"one":  1,
		"two":  2,
		"four": 4,
		"six":  6,
	}

	isOdd := func(key string, value int) bool {
		return value%2 == 1
	}

	anyOdd := gomap.Some(m, isOdd)

	fmt.Printf("Any odd: %t\n", anyOdd)
}
```

In this program, the `gomap.Some` function is called with the `m` map as its argument and a `isOdd` function that checks if the value is odd. The returned boolean is assigned to the `anyOdd` variable. The `anyOdd` variable is then printed to the console.

```bash
$ go run main.go
Any odd: true
```