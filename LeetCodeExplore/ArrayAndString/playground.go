package main

import (
	"fmt"
	"sort"
)

// Introduction to Array
func main() {
	// 1. Initialize
	//var a0 [5]int
	// other elements will be set as the default value
	a1 := [5]int{1, 2, 3}

	// 2. Get Length
	size := len(a1)
	fmt.Println("The size of a1 is:", size)

	// 3. Access Element
	fmt.Println("The first element is:", a1[0])

	// 4. Iterate all Elements
	fmt.Println("[Version 1] The contents of a1 are:")
	for i := 0; i < size; i++ {
		fmt.Printf(" ", a1[i])
	}
	fmt.Println()

	fmt.Print("[Version 2] The contents of a1 are:")
	for _, item := range a1 {
		fmt.Print(" ", item)
	}
	fmt.Println()

	// 5. Modify Element
	a1[0] = 4

	// 6. Sort
	// Convert array to slice for sorting
	a1Slice := a1[:]
	sort.Ints(a1Slice)
}
