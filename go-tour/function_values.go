package main

/*
 3.24 Function values

 Functions are values too. They can be passed around just like other values.
 Function values may be used as function arguments and return values.

*/

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypoth_fn := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypoth_fn(5, 12))

	fmt.Println(compute(hypoth_fn))
	fmt.Println(compute(math.Pow))
}

// go run function_values.go
// 13
// 5
// 81
