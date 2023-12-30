package main

import (
	"fmt"
)

func main() {
	tutorial()
}

type Inner2 struct {
	A int
}

func (i Inner2) IntPrinter(val int) string {
	return fmt.Sprintf("Inner2: %d", val)
}

func (i Inner2) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer2 struct {
	Inner2
	S string
}

func (o Outer2) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func tutorial() {
	o := Outer{
		Inner: Inner2{
			A: 10,
		},
		S: "Hello",
	}
	fmt.Println(o.Double())
}
