package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// to func Scale(v Vertex, f float64) causes an error
// standalone functions are strict about types and do not automatically dereference pointers.
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// to compile change Scale(&v, 10) to Scale(v, 10).
	// but the print shows the number 5 again.
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
