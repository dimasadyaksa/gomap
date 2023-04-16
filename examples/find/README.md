## Find

The Find function is used to find a value in a map based on a finder function
that returns a boolean. It takes a map and a finder function as its arguments.

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

	finder := func(key string, value int) bool {
		return value == 2
	}

	value, ok := gomap.Find(m, finder)
	if ok {
		fmt.Println(value)
		return
	}

	fmt.Println("not found")
}
```

In this program, the `gomap.Find` function is called with the `m` map as its argument and a `finder`
function that returns true if the value is equal to 2. If the value is found, it is assigned to the
`value` variable and the `ok` variable is set to true.

```bash
$ go run main.go
2
```
