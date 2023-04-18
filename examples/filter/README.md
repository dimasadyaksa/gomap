## Filter

The `Filter` function is used to filter a map based on a filter function that returns a boolean. 
It takes a map and a filter function as its arguments. 

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

	filterer := func(key string, value int) bool {
		return value == 2
	}

	filtered := gomap.Filter(m, filterer)

	for key, value := range filtered {
		fmt.Printf("%s: %d\n", key, value)
	}
}
```

In this program, the `gomap.Filter` function is called with the `m` map as its argument and a `filterer`
function that returns true if the value is equal to 2. The returned map is assigned to the `filtered` variable.
The `filtered` variable is then iterated over and the keys and values are printed to the console.

```bash
$ go run main.go
two: 2
```