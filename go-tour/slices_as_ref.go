package main

/*
 3.8 Slices are like references to arrays

 A slice does not store any data, it just describes a section of an underlying array.
 Changing the elements of a slice modifies the corresponding elements of its underlying array.
 Other slices that share the same underlying array will see those changes.

*/

import "fmt"

func main() {
	names := [4]string{
		"Roger",
		"David",
		"Rick",
		"Nick",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "FooBar"
	// see the changes
	fmt.Println(a, b)
	fmt.Println(names)
}

// go run slices_as_ref.go
// [Roger David Rick Nick]

// [Roger David] [David Rick]

// [Roger FooBar] [FooBar Rick]
// [Roger FooBar Rick Nick]
