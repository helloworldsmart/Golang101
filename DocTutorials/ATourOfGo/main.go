package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func main() {
	pic.Show(Pic)
	//tutorial23()
	//tutorial24()
	//tutorial25()
	tutorial26()
}

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		result[y] = make([]uint8, dx)

		for x := 0; x < dx; x++ {
			result[y][x] = uint8((x + y) / 2)
			result[y][x] = uint8(x * y)
			result[y][x] = uint8(x ^ y)
		}
	}
	return result
}

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func tutorial23() {
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.399967,
	}
	fmt.Println(m["Bell Labs"])
}

func tutorial24() {
	var m = map[string]Vertex2{
		"Bell Labs": Vertex2{
			40.68433, -74.39967,
		},
		"Google": Vertex2{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)
}

func tutorial25() {
	var m = map[string]Vertex2{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}

func tutorial26() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
