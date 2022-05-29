package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		slice := make([]uint8, dx)
		for j := range slice {
			slice[j] = uint8(i * j)
		}
		pic[i] = slice
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
