package main

import (
	"fmt"

	"github.com/Lemuriets/search_matrix_diagonales/internal/helper"
	"github.com/Lemuriets/search_matrix_diagonales/internal/matrix"
)

func main() {
	height := helper.ScanUint("Input the height of matrix: ")
	width := helper.ScanUint("Input the width of matrix: ")

	mtx := matrix.NewMatrix(height, width)

	for _, v := range mtx.GetDiagonales() {
		fmt.Println(v)
	}
}
