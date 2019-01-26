package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i <= 100; i++ {
		if math.Mod(float64(i), 3.0) == 0 {
			fmt.Print("Fizz")
		}
		if math.Mod(float64(i), 5.0) == 0 {
			fmt.Print("Buzz")
		}
		if math.Mod(float64(i), 3.0) != 0 && math.Mod(float64(i), 5.0) != 0 {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
