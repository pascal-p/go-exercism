package main

/*
 3.16 Ranges

 The range form of the for loop iterates over a slice or map.

 When ranging over a slice, two values are returned for each iteration.
 The first is the index, and the second is a copy of the element at that index.

*/

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for ix, v := range pow {
		fmt.Printf("2 ** %d = %d\n", ix, v)
	}
}

// go run range.go
// 2 ** 0 = 1
// 2 ** 1 = 2
// 2 ** 2 = 4
// 2 ** 3 = 8
// 2 ** 4 = 16
// 2 ** 5 = 32
// 2 ** 6 = 64
// 2 ** 7 = 128
