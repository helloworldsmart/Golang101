package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	tutorial7()
	tutorial8()
	tutorial11()
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

var i interface{}

func tutorial9() {
	i = 20
	i = "hello"
	i = struct {
		FirstName string
		LastName  string
	}{"Fred", "Fredson"}
}

func tutorial10() {
	data := map[string]interface{}{}
	contents, err := ioutil.ReadFile("testdata/sample.json")
	if err != nil {
		//return err
	}
	//defer contents.Close()
	json.Unmarshal(contents, &data)
}

type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}

// Not Good
func (ll *LinkedList) Insert(pos int, val interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList{
			Value: val,
			Next:  ll,
		}
	}
	//ll.Next == ll.Next.Insert(pos-1, val)
	return ll
}

type MyInt int

func tutorial11() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)
}
