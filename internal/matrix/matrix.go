package matrix

import (
	"github.com/Lemuriets/search_matrix_diagonales/internal/helper"
)

type M [][]uint

type Matrix interface {
	GetMtx() M
	GetHeight() uint
	GetWidth() uint
	GetDiagonales() M
}

type mtx struct {
	Mtx    M
	Height uint
	Width  uint
}

func NewMatrix(height, width uint) Matrix {
	var result M

	for i := 0; uint(i) < height; i++ {
		result = append(result, make([]uint, width))
	}
	return &mtx{
		Mtx:    result,
		Height: height,
		Width:  width,
	}
}

func (m *mtx) GetMtx() M {
	return m.Mtx
}

func (m *mtx) GetHeight() uint {
	return m.Height
}

func (m *mtx) GetWidth() uint {
	return m.Width
}

func (m *mtx) GetDiagonales() M {
	result := m.Mtx

	minLenght := helper.MinUint(m.Height, m.Width)
	for i := 0; uint(i) < minLenght; i++ {
		result[i][i] = 1
		result[i][minLenght-uint(i)-1] = 1
	}
	return result
}
