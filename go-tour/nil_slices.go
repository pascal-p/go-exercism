package main

/*
 3.12 Nil slices

 The zero value of a slice is nil.
 A nil slice has a length and capacity of 0 and has no underlying array.

*/

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil!")
	}
}

// go run nil_slices.go
// [] 0 0
// nil!
