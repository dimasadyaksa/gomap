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
