package main

import "fmt"

func main() {
	divMain()
	myFuncMain()
}

func divMain() {
	result := div(5, 2)
	fmt.Println(result)
}

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func myFuncMain() {
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	})
	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})
}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) {
	// Do Something
}

func tutorial() {

}
