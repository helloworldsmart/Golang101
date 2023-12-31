package main

import (
	"errors"
	"fmt"
)

func main() {
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
	tutorial14()
}

//func tutorial() {
//	var variableName type
//	var vname1, vmane2, vname3 type
//	var variableName2 type = 12
//	var vname4, vmane5, vname6 type = 13, 14, 15
//	var vname7, vmane8, vname9 = 13, 14, 15
//	var vname10, vmane11, vname12 := 13, 14, 15
//	_, b := 34, 35
//	const constantName = "Michael"
//	const Pi float64 = 3.1415926
//
//	const Pi2 = 3.1415926
//	const i = 10000
//	const MaxThread = 10
//	const prefix = "astaxie_"
//
//	var isActive bool
//	var enable, disabled = true, false
//
//	var c complex64 = 5+5i
//	fmt.Printf("Value is: %v",c)
//
//	var frenchHello string
//	var emptyString string = ""
//
//}

//func test() {
//	var available bool
//	valid := false
//	available = true
//}

//func test2() {
//	no, yes, maybe := "no", "yes", "maybe"
//	japanseHello := "Konichiwa"
//	//frenchHello = "Bonjour"
//}

func tutorial3() {
	// bad
	//var s string = "hello"
	//s[0] = 'c'

	// good
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)
}

func tutorial4() {
	s := "hello,"
	m := " world"
	a := s + m
	fmt.Println("%s\n", a)
}

func tutorial5() {
	s := "hello"
	s = "c" + s[1:]
	fmt.Printf("%s\n", s)
}

func tutorial6() {
	m := `hello
			world`
	fmt.Printf("%s\n", m)
}

func tutorial7() {
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Color string

const (
	Red   Color = "Red"
	Green Color = "Green"
	Blue  Color = "Blue"
)

func printDay(day Weekday) {
	switch day {
	case Sunday:
		fmt.Println("Sunday")
	case Monday:
		fmt.Println("Monday")
	case Tuesday:
		fmt.Println("Tuesday")
	case Wednesday:
		fmt.Println("Wednesday")
	case Thursday:
		fmt.Println("Thursday")
	case Friday:
		fmt.Println("Friday")
	case Saturday:
		fmt.Println("Saturday")
	default:
		fmt.Println("Unknown Day")
	}
}

func tutorial8() {
	printDay(Sunday)
	printDay(7)
}

func tutorial9() {
	const (
		x = iota
		y = iota
		z = iota
		w
	)

	const v = iota

	const (
		h, i, j = iota, iota, iota
	)

	const (
		a       = iota
		b       = "B"
		c       = iota
		d, e, f = iota, iota, iota
		g       = iota
	)

	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
}

func tutorial10() {
	//var arr [n]type

	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The first element is %d\n", arr[0])
	fmt.Printf("The last element is %d\n", arr[9])

	a := [3]int{1, 2, 3}
	b := [10]int{1, 2, 3}
	c := [...]int{4, 5, 6}

	fmt.Printf("%s\n", a)
	fmt.Printf("%s\n", b)
	fmt.Printf("%s\n", c)
}

func tutorial11() {
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	fmt.Println(doubleArray, easyArray)
}

func tutorial12() {
	//var slice []int
	sliceValue := []byte{'a', 'b', 'c', 'd'}
	fmt.Println(sliceValue)
}

func tutorial13() {
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var a, b []byte

	a = ar[2:5]
	b = ar[3:5]
	fmt.Println(a, b)

}

func tutorial14() {
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var aSlice, bSlice []byte

	aSlice = array[:3]
	aSlice = array[5:]
	aSlice = array[:]

	aSlice = array[3:7]
	bSlice = aSlice[1:3]
	bSlice = aSlice[:3]
	bSlice = aSlice[0:5]
	bSlice = aSlice[:]
	fmt.Println(bSlice)
}
