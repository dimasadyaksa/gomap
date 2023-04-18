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
		fmt.Printf("Key: %s, Value: %d", key, value)
	}
}
