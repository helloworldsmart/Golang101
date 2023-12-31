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
	tutorial7()
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
