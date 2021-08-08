package main

/*
 3.21 Map literals continued

 If the top-level type is just a type name, you can omit it from the elements of the literal.

*/

import "fmt"

type Vertex struct {
	Lat, Lon float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}

// go run map_literals_cont_ed.go
// map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
