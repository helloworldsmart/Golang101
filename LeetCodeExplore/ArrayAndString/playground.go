package main

import (
	"fmt"
	"sort"
)

// Introduction to Array
//func main() {
//	// 1. Initialize
//	//var a0 [5]int
//	// other elements will be set as the default value
//	a1 := [5]int{1, 2, 3}
//
//	// 2. Get Length
//	size := len(a1)
//	fmt.Println("The size of a1 is:", size)
//
//	// 3. Access Element
//	fmt.Println("The first element is:", a1[0])
//
//	// 4. Iterate all Elements
//	fmt.Println("[Version 1] The contents of a1 are:")
//	for i := 0; i < size; i++ {
//		fmt.Printf(" ", a1[i])
//	}
//	fmt.Println()
//
//	fmt.Print("[Version 2] The contents of a1 are:")
//	for _, item := range a1 {
//		fmt.Print(" ", item)
//	}
//	fmt.Println()
//
//	// 5. Modify Element
//	a1[0] = 4
//
//	// 6. Sort
//	// Convert array to slice for sorting
//	a1Slice := a1[:]
//	sort.Ints(a1Slice)
//}

// Introduction to Dynamic Array
func main() {
	// 1. Initialize
	var v0 []int
	v1 := make([]int, 5) // Initialized with zero values
	fmt.Println("v0: ", v0)

	// 2. Make a copy
	v2 := make([]int, len(v1))
	copy(v2, v1)
	v3 := append([]int(nil), v2...)
	fmt.Println("v3: ", v3)

	// 3. Cast an array to a slice
	a := [5]int{0, 1, 2, 3, 4}
	v4 := a[:]
	fmt.Println("v4: ", v4)

	// 4. Get length
	fmt.Println("The size of v4 is:", len(v4))

	// 5. Access element
	fmt.Println("The first element in v4 is:", v4[0])

	// 6. Iterate the slice
	fmt.Print("[Version 1] The content of v4 are:")
	for i := 0; i < len(v4); i++ {
		fmt.Print(" ", v4[i])
	}
	fmt.Println()

	fmt.Print("[Version 2] The contents of v4 are:")
	for _, item := range v4 {
		fmt.Print(" ", item)
	}
	fmt.Println()

	fmt.Print("[Version 3] The contents of v4 are:")
	for i := range v4 {
		fmt.Print(" ", v4[i])
	}
	fmt.Println()

	// 7. Modify element
	v4[0] = 5

	// 8. Sort
	sort.Ints(v4)

	// 9. Add New element at the end of the slice
	v4 = append(v4, -1)

	// 10. Delete the last element
	v4 = v4[:len(v4)-1]
}
