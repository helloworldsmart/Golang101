package main

import "golang.org/x/tour/pic"

func main() {
	pic.Show(Pic)
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
