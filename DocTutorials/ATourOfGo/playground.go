package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	//tutorial()
	//tutorial2()
	//tutorial3()
	//tutorial4()
	//tutorial5()
	//tutorial6()
	tutorial7()
}

func tutorial() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

const (
	// Create a huge number by shifting a 1 bit left 100 places
	// In other words, the binary number that is 1 followed by 100 zeroes
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func tutorial2() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func tutorial3() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func tutorial4() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func tutorial5() {
	fmt.Println(sqrt(2), sqrt(-4))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func tutorial6() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func customSqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		prevZ := z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %d: %v\n", i+1, z)

		if customAbs(prevZ-z) < 1e-6 {
			break
		}
	}
	return z
}

func customAbs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func tutorial7() {
	fmt.Println(customSqrt(2))
}

func tutorial8() {
	fmt.Println("The operating system is: ", runtime.GOOS)
}
