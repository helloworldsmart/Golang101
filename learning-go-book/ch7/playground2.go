package main

import (
	"fmt"
)

func main() {
	tutorial7()
	tutorial8()
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

type LogicProvider struct {
}

func (lp LogicProvider) Process(data string) string {
	// 商業邏輯
	return ""
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	// 從某處取得資料
	//c.L.Process(data)
}

func tutorial8() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Closer() error
}

type ReadCloser interface {
	Reader
	Closer
}
