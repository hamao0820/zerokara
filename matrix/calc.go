package matrix

import "gonum.org/v1/gonum/mat"

func add(a, b Matrix) Matrix {
	var dense mat.Dense
	dense.Add(a, b)
	return &matrix{Dense: &dense}
}

func sub(a, b Matrix) Matrix {
	var dense mat.Dense
	dense.Sub(a, b)
	return &matrix{Dense: &dense}
}

func mul(a, b Matrix) Matrix {
	var dense mat.Dense
	dense.MulElem(a, b)
	return &matrix{Dense: &dense}
}
func matMul(a, b Matrix) Matrix {
	var dense mat.Dense
	dense.Mul(a, b)
	return &matrix{Dense: &dense}
}

func div(a, b Matrix) Matrix {
	var dense mat.Dense
	dense.DivElem(a, b)
	return &matrix{Dense: &dense}
}

func scale(a float64, b Matrix) Matrix {
	var dense mat.Dense
	dense.Scale(a, b)
	return &matrix{Dense: &dense}
}

func is2D(matrix *matrix) bool {
	return matrix.RawMatrix().Rows > 1
}

func is1D(matrix *matrix) bool {
	return matrix.RawMatrix().Rows == 1
}

func addf(a Matrix, b float64) Matrix {
	var dense mat.Dense
	dense.Add(a, a.CopyFull(b))
	return &matrix{Dense: &dense}
}

func subf(a Matrix, b float64) Matrix {
	var dense mat.Dense
	dense.Sub(a, a.CopyFull(b))
	return &matrix{Dense: &dense}
}

func mulf(a Matrix, b float64) Matrix {
	var dense mat.Dense
	dense.MulElem(a, a.CopyFull(b))
	return &matrix{Dense: &dense}
}
func divf(a Matrix, b float64) Matrix {
	var dense mat.Dense
	dense.DivElem(a, a.CopyFull(b))
	return &matrix{Dense: &dense}
}
