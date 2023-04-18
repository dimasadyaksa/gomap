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
