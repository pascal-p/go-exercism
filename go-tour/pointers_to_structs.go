package main

/*

 3.4 Pointers to structs
 Struct fields can be accessed through a struct pointer.

 To access the field X of a struct when we have the struct pointer p we could write (*p).X.
 However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

*/

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X, p.Y = 1e9, 1e2
	fmt.Println(v)
}

// go run pointers_to_structs.go
// {1000000000 100}
