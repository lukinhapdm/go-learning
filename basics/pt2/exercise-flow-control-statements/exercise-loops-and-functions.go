package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// Start with initial guess
	z := 1.0

	// Repeat until the change is very small
	for {
		// Calculate the new guess
		newZ := z - (z*z-x)/(2*z)
		fmt.Printf("Current guess: %.15f\n", z)

		// Check if the change is negligible
		if math.Abs(newZ-z) < 1e-15 {
			return newZ
		}

		z = newZ
	}
}

func main() {
	numbers := []float64{2, 20, 10}

	for _, n := range numbers {
		fmt.Println("Final result:", Sqrt(n))
		fmt.Println("Actual sqrt(2):", math.Sqrt(n))
	}
}
