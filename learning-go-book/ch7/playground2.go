package main

import (
	"fmt"
)

func main() {
	tutorial7()
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

type Outer3 struct {
	Inner2
	S string
}

func (o Outer3) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func tutorial7() {
	o := Outer3{
		Inner2: Inner2{
			A: 10,
		},
		S: "Hello",
	}
	fmt.Println(o.Double())
}
