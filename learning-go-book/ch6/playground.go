package main

import "fmt"

func main() {
	//tutorial2()
	tutorial4()
}

func tutorial() {
	//var x int32 = 10
	//var y bool = true
	//pointerX := &x
	//pointerY := &y
	//var pointerZ *string
}

func tutorial2() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)
	z := 5 + *pointerToX
	fmt.Println(z)
}

func tutorial3() {
	var x *int
	fmt.Println(x == nil)
	fmt.Println(*x)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1086e98]
}

func tutorial4() {
	x := 10
	var pointerToX *int
	pointerToX = &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)
}
