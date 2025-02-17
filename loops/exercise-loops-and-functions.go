package main

import (
	"fmt"
	"math"
)

/** Implement this in the func Sqrt provided. A decent starting guess for z is 1,
 * no matter what the input. To begin with, repeat the calculation 10 times and
 * print each z along the way. See how close you get to the answer for various values
 * of x (1, 2, 3, ...) and how quickly the guess improves.
 */
func Sqrt10iterations(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

/**
 * Next, change the loop condition to stop once the value has stopped changing
 * (or only changes by a very small amount). See if that's more or fewer than
 * 10 iterations. Try other initial guesses for z, like x, or x/2. How close are
 * your function's results to the math.Sqrt in the standard library?
 */
func Sqrt(x float64) float64 {
	z := x / 2
	threshold := 1e-10

	for {
		newZ := z - (z*z-x)/(2*z)
		if math.Abs(newZ-z) < threshold {
			break
		}
		z = newZ
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
