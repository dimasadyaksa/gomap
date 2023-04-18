package main

import (
	"fmt"

	"github.com/dimasadyaksa/gomap"
)

type Student struct {
	Name  string
	Score int
}

func main() {
	classA := map[int]Student{
		1: {"John", 100},
		2: {"Bob", 90},
		3: {"Alice", 80},
	}

	classB := map[int]Student{
		1: {"Matthew", 100},
		2: {"Andrew", 90},
		3: {"Doe", 80},
	}

	isScoreEqual := func(a, b Student) bool {
		return a.Score == b.Score
	}

	isEqual := gomap.EqualFunc(classA, classB, isScoreEqual)

	fmt.Printf("Is Equal: %t\n", isEqual)
}
