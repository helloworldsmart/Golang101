package main

import (
	"fmt"
	"time"
)

func main() {
	//tutorial()
	//tutorial1()
	//tutorial2()
	//tutorial3()
	//tutorial4()
	//tutorial5()
	tutorial6()
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Score int
type Converter func(string) Score
type TeamScores map[string]Score

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func tutorial() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}

func tutorial1() {
	var c Counter
	doUpdateWrong(c)
	fmt.Println("in main:", c.String())
	doUpdateRight(&c)
	fmt.Println("in main:", c.String())
}

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

func tutorial2() {
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(12)) // false
}

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func tutorial3() {
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))

	// 方法值
	f1 := myAdder.AddTo
	fmt.Println(f1(10))

	// 方法表達值
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 15))
}

type HighScore Score
type Employee Person

func tutorial4() {
	var i int = 300
	var s Score = 100
	var hs HighScore = 200
	// hs = s // Cannot use 's' (type Score) as the type HighScore
	//s = i
	s = Score(i)
	hs = HighScore(s)
	fmt.Println(hs)
}

type Employee2 struct {
	Name string
	ID   string
}

func (e Employee2) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee2
	Reports []Employee2
}

func (m Manager) FindNewEmployee() []Employee2 {
	// Do something
	return []Employee2{}
}

func tutorial5() {
	m := Manager{
		Employee2: Employee2{
			Name: "Bob Bobson",
			ID:   "123456",
		},
		Reports: []Employee2{},
	}

	fmt.Println(m.ID)
	fmt.Println(m.Description())
}

type Inner struct {
	X int
}

type Outer struct {
	Inner
	X int
}

func tutorial6() {
	o := Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.X)
	fmt.Println(o.Inner.X)
}
