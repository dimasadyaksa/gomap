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
