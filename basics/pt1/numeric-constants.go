package main

import "fmt"

const (
	Big   = 1 << 100
	Small = big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(big))

	/*
		The compiler rejects since exceeds the maximum value an int can hold
		fmt.Println(needInt(Big))
	*/
}
