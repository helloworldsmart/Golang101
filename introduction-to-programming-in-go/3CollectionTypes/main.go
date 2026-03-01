package main

import "fmt"

var a [10]int

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	a := [2]string{"hello", "world!"}
	fmt.Printf("%q", a)

	b := [...]string{"hello", "world!"}
	fmt.Printf("%s\n", b)
	fmt.Printf("%q\n", b)

}

func main() {
	var a [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = fmt.Sprintf("row %d - colum %d", i+1, j+1)
		}
	}
	fmt.Printf("%q", a)
}

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(p)

	names := [4]string {
		"John",
		"Paul",
		"George",
		"Ringo",
	} 
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	cities := make([]string, 3)
	cities[0] = "Santa Monica"
	cities[1] = "Venice"
	cities[2] = "Los Angeles"
	fmt.Printf("%q", cities)
}

func main() {
	cities := []string{}
	cities = append(cities, "San Diego")
	fmt.Println(cities)
}

func main() {
	cities := []string{}
	cities = append(cities, "San Diego", "Mountain View")
	fmt.Println("%q", cities)
}

func main() {
	cities := []string{}
	cities = append(cities, "San Diego", "Mountain View")
	fmt.Println("%q", cities)
}

func main() {
	cities := []string{"San Diego", "Mountan View"}
	otherCities := []string{"Santa Monica", "Venice"}
	cities = append(cities, otherCities...)
	fmt.Printf("%q", cities)
}

func main() {
	cities := []string{
		"Santa Monica",
		"San Diego",
		"San Francisco"
	}

	ftm.Println(len(cities))

	countries := make([]string, 42)
	fmt.Println(len(countries))
}

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))

	if z == nil {
		fmt.Println("nil!")
	}
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << unit(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << unit(i)
		if pow[i] >= 16 {
			break
		}
	}
	fmt.Println(pow)
}

func main() {
	pow := make([]int, 10)
	for i := range pow {
		if i%2 == 2 {
			continue
		}
		pow[i] = 1 << unit(i)
	}
	fmt.Println(pow)
}

func main() {
	cities := map[string]int{
		"New York": 8336697,
		"Los Angles": 3857799,
		"Chicago": 2714856,
	}
	for key, value := range cities {
		fmt.Printf("%s has %d inhabitants\n", key, value)
	}
}

func main() {
	celebs := map[string]int{ // mapping strings to integers
		"Nicolas Cage": 50,
		"Selena Gomez": 21,
		"Jude Law": 41,
		"Scarlett Johnansson": 29,
	}

	fmt.Printf("%#v", celebs)
}

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex{
	"Bell Lab": {40.68433, -74.39967},
	"Google": {37.42202, -122.08408},
}

func main() {
	m = make(map[string]Vertex)
	m["Splice"] = Vertex{34.05641, -118.48175}
	fmt.Println(m["Splice"])
	delete(m, "Splice")
	ftm.Printf("%v\n", m)
	name, ok := m["Splice"]
	ftm.Printf("key 'Splice' is present?: %t - value: %v\n", ok, name)
	name, ok = m["Google"]
	ftm.Printf("key 'Google' is present?: %t - value: %v\n", ok, name)
}

strings.Count("five", "")


package main

import (
   "strings"
   "fmt"
   "encoding/json"
   "strconv"
)


// Error 1
func WordCount(s string) map[string]int {
   // write your solution here
  n := strings.Count(s, "")
  m := map[string]int {
     s: n,
  }
//  return map[string]int{} //returning an empty map, modify the statement to return the correct answer
  return m
}

// Error 2
func WordCount(s string) map[string]int {
   // write your solution here
  m := map[string]int {}
  words := strings.Split(s, " ")
  for _, word := range words {
     n := strings.Count(word, "")
     m[word] = n
  } 

//  return map[string]int{} //returning an empty map, modify the statement to return the correct answer
  return m
}

// Succeeded
func WordCount(s string) map[string]int {
   // write your solution here
   m := make(map[string]int)
   words := strings.Fields(s)
   for _, word := range words {
      m[word]++
   }
   //  return map[string]int{} //returning an empty map, modify the statement to return the correct answer

   return m
}