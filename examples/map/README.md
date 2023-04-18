## Map

The `Map` function is used to map a map based on a mapper function. It takes a map and a mapper function as its arguments.

```go
package main

import (
	"fmt"
	"github.com/dimasadyaksa/gomap"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}

	mapper := func(i int) int {
		return slice[i]
	}

	m := gomap.Map(slice, mapper)

	for key, value := range m {
		fmt.Printf("Key: %d, Value: %d\n", key, value)
	}

}
```

In this program, the `gomap.Map` function is called with the `slice` 
slice as its argument and a `mapper` function that returns the value of the slice at the index that will be used as the map key.
The returned map is assigned to the `m` variable. The `m` variable is then iterated over and the keys and values are printed to the console.

```bash
$ go run main.go
Key: 1, Value: 1
Key: 2, Value: 2
Key: 3, Value: 3
Key: 4, Value: 4
Key: 5, Value: 5
```