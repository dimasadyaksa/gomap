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

	values := gomap.Values(m)

	for _, value := range values {
		fmt.Println(value)
	}
}
