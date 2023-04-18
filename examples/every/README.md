## Every

The `Every` function is used to check if all values in a map satisfy a predicate function. It takes a map and a predicate function as its arguments.

```go
package main

import (
	"fmt"
	"github.com/dimasadyaksa/gomap"
)

func main() {
	m := map[string]int{
		"two":   2,
		"four":  4,
		"six":   6,
		"eight": 8,
	}

	isEven := func(key string, value int) bool {
		return value%2 == 0
	}

	allEven := gomap.Every(m, isEven)

	fmt.Printf("All even: %t\n", allEven)
}
```

In this program, the `gomap.Every` function is called with the `m` map as its argument and a `isEven` function that checks if the value is even. The returned boolean is assigned to the `allEven` variable. The `allEven` variable is then printed to the console.

```bash
$ go run main.go
All even: true
```