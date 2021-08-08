package main

/*
 3.5 Struct Literals

 A struct literal denotes a newly allocated struct value by listing the values of its fields.
 You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
 The special prefix & returns a pointer to the struct value.

*/

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implict
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex (pointer to Vertex)
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

// go run struct_literals.go
// {1 2} &{1 2} {1 0} {0 0}
