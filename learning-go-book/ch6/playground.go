package main

import "fmt"

func main() {
	//tutorial2()
	//tutorial4()
	//tutorial5()
	//tutorial6()
	tutorial7()
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

func tutorial5() {
	var x = new(int)
	fmt.Println(x == nil)
	fmt.Println(*x)
}

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func tutorial6() {
	p := person{
		FirstName: "Pat",
		//MiddleName: "Perry", // '"Perry"' (type string) cannot be represented by the type *string
		//MiddleName: &"Perry", // Cannot take the address of '"Perry"'
		MiddleName: stringp("Perry"),
		LastName:   "Peterson",
	}
	fmt.Println(p)
}

func stringp(s string) *string {
	return &s
}

func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

func tutorial7() {
	x := 10
	failedUpdate(&x)
	fmt.Println(x)
	update(&x)
	fmt.Println(x)
}
