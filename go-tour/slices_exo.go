package main

/*
 3.18 Exercise: Slices

 Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
 When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

 The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

 (You need to use a loop to allocate each []uint8 inside the [][]uint8.)
 (Use uint8(intValue) to convert between types.)

*/

import "golang.org/x/tour/pic" // "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {

	pic := make([][]uint8, dy)

	for y := range pic {

		pic[y] = make([]uint8, dx)

		for x := range pic[y] {
			pic[y][x] = uint8((x + y) / 2)
		}
	}

	return pic
}

/*
func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy) // [dy][]uint

	for ix := 0; ix < dy; ix++ {
		pic[ix] = make([]uint8, dx) // [dx]uint8
	}

	for ix := range pic {
		for jx := range pic[ix] {
			switch {
			case jx%15 == 0:
				pic[ix][jx] = 240

			case jx%3 == 0:
				pic[ix][jx] = 120

			case jx%5 == 0:
				pic[ix][jx] = 150

			default:
				pic[ix][jx] = 100
			}
		}
	}

	return pic
}
*/

func main() {
	Pic.show(Pic)
}
