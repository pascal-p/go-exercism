package main

/*
 3.11 Slice length and capacity

 A slice has both a length and a capacity.
 The length of a slice is the number of elements it contains.

 The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

 The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

 You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
 Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

*/

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length
	s = s[:4]
	printSlice(s)

	// Drop its first two values
	s = s[2:]
	printSlice(s)

	// Go beyond its capacity! => Exception
	// s = s[:8]
	// printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// go run slice_len_cap.go
// len=6 cap=6 [2 3 5 7 11 13]
// len=0 cap=6 []
// len=4 cap=6 [2 3 5 7]
// len=2 cap=4 [5 7]

// panic: runtime error: slice bounds out of range [:8] with capacity 4

// goroutine 1 [running]:
// main.main()
//         /home/pascal/Projects/Exercism/go/go-tour/slice_len_cap.go:37 +0xf8
// exit status 2
