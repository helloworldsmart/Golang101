package main

import (
	"fmt"
	"math/rand"
)

func main() {
	shadowingVariable()
}

func sampleCode() {
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}

// 4-1
func shadowingVariable() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}

func forTutorial() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 1
	for j < 100 {
		fmt.Println(j)
		j = j * 2
	}

	// for {
	// 	fmt.Println("Hello")
	// }

	// do {
	// 	// do thing
	// } while (CONDITION)

	// for {
	// 	if !CONDITION {
	// 		break
	// 	}
	// }

	for k := 1; k <= 100; k++ {
		if k%3 == 0 && k%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if k%3 == 0 {
			fmt.Println("Fizz")
			continue
		}

		if k%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(k)
	}

	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}

	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}
