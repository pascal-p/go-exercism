package main

/*

 3.6 Arrays

 The type [n]T is an array of n values of type T.
 The expression
	var a [10]int
 declares a variable a as an array of ten integers.

 An array's length is part of its type, so arrays cannot be resized.
 This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
*/

import "fmt"

func main() {
	var ary [2]string
	ary[0] = "Hello"
	ary[1] = "World"
	fmt.Println(ary[0], ary[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// go run arrays.go
// Hello World
// [2 3 5 7 11 13]
