package image

import (
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"strings"
)

type Image struct {
	Layers [][][]int
	Width  int
	Height int
}

const (
	White       = 0
	Black       = 1
	Transparent = 2
)

func New(data string, width int, height int) *Image {
	var layers [][][]int
	var layer [][]int
	var row []int

	for i := 0; i < len(data); i++ {
		pixel := convert.RuneToInt(data[i])
		row = append(row, pixel)

		if (i+1)%width == 0 && row != nil {
			layer = append(layer, row)
			row = make([]int, 0)
		}

		if (i+1)%(width*height) == 0 && layer != nil {
			layers = append(layers, layer)
			layer = make([][]int, 0)
		}
	}

	return &Image{layers, width, height}
}

func (img *Image) Validate() int {
	minZeroes := ints.MaxInt
	result := 0

	for _, layer := range img.Layers {
		counts := map[int]int{
			0: 0,
			1: 0,
			2: 0,
		}

		for _, row := range layer {
			for _, pixel := range row {
				if _, ok := counts[pixel]; ok {
					counts[pixel]++
				}
			}
		}

		if counts[0] < minZeroes {
			minZeroes = counts[0]
			result = counts[1] * counts[2]
		}
	}

	return result
}

func (img *Image) Decode() [][]int {
	result := make([][]int, img.Height)

	for h := 0; h < img.Height; h++ {
		row := make([]int, img.Width)
		for w := 0; w < img.Width; w++ {
			color := Transparent
			for i := 0; i < len(img.Layers) && color == Transparent; i++ {
				color = img.Layers[i][h][w]
			}
			row[w] = color
		}
		result[h] = row
	}

	return result
}

func Draw(image [][]int) string {
	result := ""

	for _, row := range image {
		line := make([]string, len(row))
		for col, value := range row {
			if value == Black {
				line[col] = "X"
			} else {
				line[col] = " "
			}
		}

		result += strings.Join(line, "")
		result += "\n"
	}

	return result
}
