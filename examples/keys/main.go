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
