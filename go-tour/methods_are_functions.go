package main

/*
 4.2 Methods are Functions

 Remember: a method is just a function with a receiver argument.
 Here's Abs written as a regular function with no change in functionality.
*/

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
