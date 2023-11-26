package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//shadowingVariable()
	//shadowingVariable2()
	//shadowingVariable3()
	//ifElseSampleCode()
	//ifElseSampleCode2()
	//forSampleCode()
	//forSampleCode2()
	//breakContinue2()
	//breakContinue3()
	//forRange()
	//forRange2()
	//forRange3()
	//forRange4()
	forRange5()
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

// 4 - 2
func shadowingVariable2() {
	x := 10
	if x > 5 {
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)
}

// 4 - 3 : fmt.Println undefined (type string has no field or method Println)
func shadowingVariable3() {
	x := 10
	fmt.Println(x)
	//fmt := "oops"
	//fmt.Println(fmt)
}

// 4 - 5
func ifElseSampleCode() {
	n := rand.Intn(10)
	fmt.Println("Number:", n)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's tpp big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}

// 4 - 6
func ifElseSampleCode2() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's good number:", n)
	}
	// 4 - 7 : undefined: n
	//fmt.Println(n)
}

// 4 - 8
func forSampleCode() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// 4 - 9
func forSampleCode2() {
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
}

// 4 - 10
func forSampleCode3() {
	for {
		fmt.Println("Hello")
	}
}

func breakContinue() {
	// JAVA
	//do {
	//	// handle action
	//} while (CONDITION)

	// GO version
	//for {
	//	if !CONDITION {
	//		break
	//	}
	//}
}

// 4 - 11
func breakContinue2() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// 4 - 12
func breakContinue3() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}

// 4 - 13
func forRange() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}

// 4 - 14
func forRange2() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		fmt.Println(v)
	}
}

func forRange3() {
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}
}

// 4 - 15
func forRange4() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}

func forRange5() {
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
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
