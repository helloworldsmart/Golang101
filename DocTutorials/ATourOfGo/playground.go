package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	//tutorial()
	//tutorial2()
	//tutorial3()
	//tutorial4()
	//tutorial5()
	//tutorial6()
	//tutorial7()
	//tutorial8()
	//tutorial9()
	//tutorial10()
	//tutorial11()
	//tutorial12()
	//tutorial13()
	//tutorial14()
	//tutorial15()
	//tutorial16()
	//tutorial17()
	tutorial18()
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
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsb, openbsb
		// plan9, windows
		fmt.Printf("%s.\n", os)
	}
}

func tutorial9() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away")
	}
}

func tutorial10() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening")
	}
}

func tutorial11() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func tutorial12() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}

func tutorial13() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j // point to j
	fmt.Println(*p)
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of 1
}

type Vertex struct {
	X int
	Y int
}

func tutorial14() {
	fmt.Println(Vertex{1, 2})
}

func tutorial15() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func tutorial16() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

var (
	v1 = Vertex{1, 2} // has type Vertex
	p  = &Vertex{1, 2}
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{}
)

func tutorial17() {
	fmt.Println(v1, v2, v3, p)
}

func tutorial18() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)
}

// Slices of slice

func tutorial19() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
}
