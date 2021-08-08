package main

/*
 3.26 Exercise: Fibonacci closure

 Let's have some fun with functions.
 Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

*/

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	s, x := 0, 1
	return func() int {
		old_x := x
		x += s
		s = old_x
		return s
	}
}

func main() {
	f := fibonacci()

	for ix := 0; ix < 10; ix++ {
		fmt.Println(f())
	}
}
