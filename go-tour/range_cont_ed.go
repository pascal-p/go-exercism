package main

/*
 3.17 Range continued

 You can skip the index or value by assigning to _.

 for i, _ := range pow
 for _, value := range pow

 If you only want the index, you can omit the second variable.
 for i := range pow

*/

import "fmt"

func main() {
	pow := make([]int, 10)
	for ix := range pow {
		pow[ix] = 1 << uint(ix) // == 2 ** ix, faster
	}

	for _, val := range pow {
		fmt.Printf("%d\n", val)
	}
}

// go run range_cont_ed.go
// 1
// 2
// 4
// 8
// 16
// 32
// 64
// 128
// 256
// 512
