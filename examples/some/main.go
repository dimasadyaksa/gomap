package main

import (
	"fmt"

	"github.com/dimasadyaksa/gomap"
)

func main() {
	m := map[string]int{
		"one":  2,
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
