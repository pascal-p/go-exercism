package main

/*
 1.2 For continued

 The init and post statements are optional.

*/

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// go run for_cont_ed.go
// 1024
