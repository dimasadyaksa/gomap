## Keys

The Keys function is used to get the keys of a map. 
It takes a map as its argument and returns a slice of the keys.

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

	keys := gomap.Keys(m)

	for _, key := range keys {
		fmt.Println(key)
	}
}

```

In this program, the `gomap.Keys` function is called with the `m` map 
as its argument, and the returned slice of keys is assigned to the `keys` variable.
The `keys` variable is then iterated over and the keys are printed to the console.

```bash
$ go run main.go
one
two
three
```

The order of the keys in the slice is not guaranteed to be the same.