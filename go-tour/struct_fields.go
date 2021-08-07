package main

/*
  3.3 Struct Fields
  Struct fields are accessed using a dot.

*/

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X, v.Y = 4, 2
	fmt.Println(v.X, " ", v.Y)
}

// go run struct_fields.go
// 4   2
