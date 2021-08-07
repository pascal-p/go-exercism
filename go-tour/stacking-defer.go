package main

/*
  13: Stacking Defer
  Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
  To learn more about defer statements read this blog post: https://blog.golang.org/defer-panic-and-recover

*/

import "fmt"

func main() {
	fmt.Println("Counting")

	for ix := 0; ix < 10; ix++ {
		defer fmt.Println(ix)
	}

	fmt.Println("Done")
}

// go run stacking-defer.go
// Counting
// Done
// 9
// 8
// 7
// 6
// 5
// 4
// 3
// 2
// 1
// 0
