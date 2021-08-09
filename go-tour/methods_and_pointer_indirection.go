package main

/*
 4.6 Methods and pointer indirection

 Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

 var v Vertex
 ScaleFunc(v, 5)  // Compile error!
 ScaleFunc(&v, 5) // OK

 while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

 var v Vertex
 v.Scale(5)  // OK

 p := &v
 p.Scale(10) // OK

 For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.

*/

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func ScaleFn(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2) // == (&v).Scale(2)
	ScaleFn(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFn(p, 8)

	fmt.Println(v, p)
}

/*
 go run methods_and_pointer_indirection.go
 {60 80} &{96 72}
*/
