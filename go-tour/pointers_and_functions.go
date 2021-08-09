package main

/*
 4.5 Pointers and functions

 Here we see the Abs and Scale methods rewritten as functions.

 Again, try removing the * from line 16. Can you see why the behavior changes?
 What else did you need to change for the example to compile?

 (If you're not sure, continue to the next page.)

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

func Scale(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main_alt() {
	v := &Vertex{3, 4}
	Scale(v, 10)
	fmt.Println(Abs(*v))
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
