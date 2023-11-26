package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	//divMain()
	//myFuncMain()
	//addToMain()
	divAndRemainderMain()
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

func addToMain() {
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func divAndRemainderMain() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}

func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func divAndRemainder2(numerator, denominator int) (result int, remainder int, err error) {

	result, remainder = 20, 30
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

// Not Good
func divAndRemainder3(numerator, denominator int) (result int, remainder int, err error) {

	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result, remainder = numerator/denominator, numerator%denominator
	return
}

func tutorial() {

}
