package main

import "fmt"

func main() {
	var x = [3]int{10, 20, 30}
	var y = [12]int{1, 5: 4, 6, 10: 100, 15}
	var z = [...]int{10, 20, 30}
	var i = [...]int{1, 2, 3}
	var j = [3]int{1, 2, 3}

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(i == j)
	var k [2][3]int
	fmt.Println(k)

	fmt.Println()
}

func slice() {
	var x = []int{1, 2, 3}
	x = append(x, 4)
	x = append(x, 5, 6, 7)
	y := []int{20, 30, 40}
	x = append(x, y...)
}

func make() {
	// x := make([]int, 0, 10)
	// x = append(x, 5, 6, 7, 8)
}

func cuttingSlice() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}

func copy() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)
}
