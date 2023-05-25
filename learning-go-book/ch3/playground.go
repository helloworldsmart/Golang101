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

func mapFunc() {
	// var nilMap map[string]int
	totalWinss := map[string]int{}

	teams := map[string][]string{
		"Orcas": []string{"Fred", "Ralph", "Bijou"},
		"Lions": []string{"Sarah", "Peter", "Billie"},
	}
	ages := make(map[int][]string, 10)
}

func okMap() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	k, ok := m["goodbye"]
	fmt.Println(k, ok)

	delete(m, "hello")
}

func set() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}

func structTest() {
	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person

	bob := person{}

	julia := person{
		"Julia",
		40,
		"cat",
	}

	bob.name = "Bob"
	fmt.Println(julia.name)
}

func anonymousStruct() {
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "bob"
	person.age = 50
	person.pet = "dog"

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
}
