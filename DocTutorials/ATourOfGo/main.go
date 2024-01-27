package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func main() {
	pic.Show(Pic)
	tutorial23()
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
