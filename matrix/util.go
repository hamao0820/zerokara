package matrix

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func random(min, max float64) float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float64()*(max-min) + min
}

func RandomPermutation(length int) []int {
	result := make([]int, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, i)
	}

	rand.Shuffle(length, func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return result
}

func Allclose(a, b Matrix, rtol, atol *float64) bool {
	if rtol == nil {
		rtol = new(float64)
		*rtol = 1e-5
	}
	if atol == nil {
		atol = new(float64)
		*atol = 1e-8
	}

	for i := 0; i < b.Rows(); i++ {
		for j := 0; j < b.Cols(); j++ {
			aij := b.At(i, j)
			bij := a.At(i, j)
			if math.Abs(aij-bij) > (*atol + (*rtol)*math.Abs(bij)) {
				return false
			}
		}
	}
	return true
}

func BroadcastForMatrix(a, b Matrix) (Matrix, Matrix) {
	acol, arow := a.Shape().C, a.Shape().R
	bcol, brow := b.Shape().C, b.Shape().R
	if acol == bcol && arow == brow {
		return a, b
	}
	if acol == 1 && brow == 1 {
		newMat1 := NewMat(Shape{R: arow, C: bcol})
		newMat2 := NewMat(Shape{R: arow, C: bcol})
		return newMat1.CopyApplyWithIndex(func(i, j int, v float64) float64 { return a.At(i, 0) }),
			newMat2.CopyApplyWithIndex(func(i, j int, v float64) float64 { return b.At(0, j) })
	}
	if arow == 1 && bcol == 1 {
		newMat1 := NewMat(Shape{R: brow, C: acol})
		newMat2 := NewMat(Shape{R: brow, C: acol})
		return newMat1.CopyApplyWithIndex(func(i, j int, v float64) float64 { return a.At(0, j) }),
			newMat2.CopyApplyWithIndex(func(i, j int, v float64) float64 { return b.At(i, 0) })
	}
	if acol == 1 && arow == 1 {
		return b.CopyApply(func(f float64) float64 { return a.At(0, 0) }), b
	}
	if bcol == 1 && brow == 1 {
		return a, a.CopyApply(func(f float64) float64 { return b.At(0, 0) })
	}
	if acol == 1 {
		return b.CopyApplyWithIndex(func(i, j int, v float64) float64 { return a.At(i, 0) }), b
	}
	if bcol == 1 {
		return a, a.CopyApplyWithIndex(func(i, j int, v float64) float64 { return b.At(i, 0) })
	}
	if arow == 1 {
		return b.CopyApplyWithIndex(func(i, j int, v float64) float64 { return a.At(0, j) }), b
	}
	if brow == 1 {
		return a, a.CopyApplyWithIndex(func(i, j int, v float64) float64 { return b.At(0, j) })
	}

	panic(fmt.Sprintf("BroadcastForMatrix shape1: %s, shape2: %s", a.Shape(), b.Shape()))
}

func NumericalGradient(f func(Matrix) Matrix, x Matrix) Matrix {
	eps := 1e-4
	grad := NewMat(x.Shape())

	for i := 0; i < x.Rows(); i++ {
		for j := 0; j < x.Cols(); j++ {
			tmp := x.At(i, j)

			x.Set(i, j, tmp+eps)
			y1 := f(x)

			x.Set(i, j, tmp-eps)
			y2 := f(x)

			diff := y1.CopySub(y2).Sum().At(0, 0)
			grad.Set(i, j, diff/(2*eps))
			x.Set(i, j, tmp)
		}
	}

	return grad
}

func Merge1DMatrixes(mats ...Matrix) Matrix {
	result := [][]float64{}
	for _, m := range mats {
		if m.Shape().R > 1 {
			panic("matrix rows > 1")
		}
		result = append(result, m.Flatten())
	}
	return New2D(result)
}

func ARange(min, max, step int) []int {
	if step == 0 {
		step = 1
	}

	if min > max && step > 0 {
		panic(fmt.Sprintf("missing step: min %d, max %d, step %d", min, max, step))
	} else if step < 0 {
		panic(fmt.Sprintf("missing step: min %d, max %d, step %d", min, max, step))
	} else if min == max {
		return []int{min}
	}

	result := []int{}
	for i := min; i <= max; i += step {
		result = append(result, i)
	}

	return result
}
