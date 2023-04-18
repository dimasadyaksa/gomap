## EqualFunc

The `EqualFunc` function is used to compare two maps. It takes two maps and a function that compares two values as its arguments and returns a boolean.

```go
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
		1: {"John", 100},
		2: {"Bob", 90},
		3: {"Alice", 80},
	}

	isScoreEqual := func(a, b Student) bool {
		return a.Score == b.Score
	}

	isEqual := gomap.EqualFunc(classA, classB, isScoreEqual)

	fmt.Printf("Is Equal: %t\n", isEqual)
}
```

In this program, the `gomap.EqualFunc` function is called with the `classA` and `classB` maps as its arguments and a `isScoreEqual` function that compares the `Score` field of the `Student` struct. The returned boolean is assigned to the `isEqual` variable. The `isEqual` variable is then printed to the console.

```bash
$ go run main.go
Is Equal: true
```