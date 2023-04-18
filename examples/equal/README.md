## Equal

The `Equal` function is used to compare two maps. It takes two maps as its arguments and returns a boolean.

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
		"one":   1,
		"two":   2,
		"three": 3,
	}

	isEqual := gomap.Equal(m1, m2)

	fmt.Printf("Is Equal: %t\n", isEqual)
}
```

In this program, the `gomap.Equal` function is called with the `m1` and `m2` maps as its arguments. The returned boolean is assigned to the `isEqual` variable. The `isEqual` variable is then printed to the console.

```bash
$ go run main.go
Is Equal: true
```