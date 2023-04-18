## Combine

The `Combine` function is used to combine two maps. It takes two maps as its arguments and returns a new map.

```go
package main

import (
	"fmt"
	"github.com/dimasadyaksa/gomap"
)

func main() {
	m1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	m2 := map[string]int{
		"four": 4,
		"five": 5,
		"six":  6,
	}

	combined := gomap.Combine(m1, m2)
	for key, value := range combined {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
```

In this program, the `gomap.Combine` function is called with the `m1` and `m2` maps as its arguments. The returned map is assigned to the `combined` variable. The `combined` variable is then iterated over and the keys and values are printed to the console.

```bash
$ go run main.go
Key: one, Value: 1
Key: two, Value: 2
Key: three, Value: 3
Key: four, Value: 4
Key: five, Value: 5
Key: six, Value: 6
```
