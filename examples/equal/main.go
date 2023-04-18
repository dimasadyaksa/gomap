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
