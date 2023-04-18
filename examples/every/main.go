package main

import (
	"fmt"

	"github.com/dimasadyaksa/gomap"
)

func main() {
	m := map[string]int{
		"two":   2,
		"four":  4,
		"six":   6,
		"eight": 8,
	}

	isEven := func(key string, value int) bool {
		return value%2 == 0
	}

	allEven := gomap.Every(m, isEven)

	fmt.Printf("All even: %t\n", allEven)
}
