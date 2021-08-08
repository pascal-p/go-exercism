package main

/*
 3.25 Function closures

 Go functions may be closures. A closure is a function value that references variables from outside its body.
 The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

 For example, the adder function returns a closure. Each closure is bound to its own sum variable.

*/

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos_fn, neg_fn := adder(), adder()
	for ix := 0; ix < 10; ix++ {
		fmt.Println(
			pos_fn(ix),
			neg_fn(-2*ix),
		)
	}
}

// go run function_closures.go
// 0 0
// 1 -2
// 3 -6
// 6 -12
// 10 -20
// 15 -30
// 21 -42
// 28 -56
// 36 -72
// 45 -90
