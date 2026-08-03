package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	/*
		Without explicit type conversions the code will not compile.
		var f float64 = math.Sqrt(x*x + y*y)
		var z uint = f
	*/

	fmt.Println(x, y, z)
}
