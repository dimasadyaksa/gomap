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

	sumReducer := func(acc int, key string, value int) int {
		return acc + value
	}

	sum := gomap.Reduce(m, 0, sumReducer)

	fmt.Println(sum)
}
