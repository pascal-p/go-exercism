package main

/*
 3.20 Map Literals

 Map literals are like struct literals, but the keys are required.

*/

import "fmt"

type Vertex struct {
	Lat, Lon float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}

// go run map_literals.go
//
// map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
