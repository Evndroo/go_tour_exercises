package geometry

import (
	"golang.org/x/tour/pic"
)

func picture(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)

	for xIndex := range slice {
		sliceValue := make([]uint8, dy)

		for yIndex := range sliceValue {
			sliceValue[yIndex] = uint8((xIndex * yIndex) / 5)
		}

		slice[xIndex] = sliceValue
	}
	return slice
}

func Slices() {
	pic.Show(picture)
}
