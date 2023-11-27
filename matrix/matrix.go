package matrix

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type (
	Shape struct{ R, C int }
)

func (s Shape) String() string { return fmt.Sprintf("(rows: %d, cols: %d)", s.R, s.C) }

type (
	Matrix interface {
		mat.Matrix
		fmt.Stringer

		At(i, j int) float64
		Set(i, j int, v float64)
		Cat(index interface{}) Matrix
		SetRow(idx int, v []float64)

		Len() int
		Rows() int
		Cols() int
		Shape() Shape
		Dims() (r, c int)
		Flatten() []float64
		Clip(min, max float64) Matrix
		Search(x interface{}, y interface{}) []float64
		Array() [][]float64
		IsEmpty() bool

		Log()
		Shuffle()
		Full(float64)
		Scale(float64)
		Reshape(Shape)
		Broadcast(Shape) Matrix
		Swap(i1, j1, i2, j2 int)
		Apply(func(float64) float64)
		Each(func(i, j int, v float64))
		EachR(func(i int, v []float64))
		EachC(func(i int, v []float64))
		Sum(options ...Option) Matrix
		Max(options ...Option) Matrix
		ArgMax(options ...Option) Matrix
		Mean(options ...Option) Matrix
		Allclose(a Matrix, rtol, atol *float64) bool
		OnOff(fn func(i, j int, v float64) bool) Matrix
		ApplyWithIndex(fn func(i, j int, v float64) float64)

		Copy() Matrix
		CopyT() Matrix
		CopyLog() Matrix
		CopyAdd(v Matrix) Matrix
		CopySub(v Matrix) Matrix
		CopyMul(v Matrix) Matrix
		CopyDiv(v Matrix) Matrix
		CopyFull(float64) Matrix
		CopyScale(float64) Matrix
		CopyMatMul(Matrix) Matrix
		CopyReshape(Shape) Matrix
		CopyAddFloat(float64) Matrix
		CopySubFloat(float64) Matrix
		CopyMulFloat(float64) Matrix
		CopyDivFloat(float64) Matrix
		CopyClip(min, max float64) Matrix
		CopyApply(func(float64) float64) Matrix
		CopyApplyWithIndex(func(i, j int, v float64) float64) Matrix
	}
	randSeedOption struct{ seed *float64 }
	RandSeedOption func(*randSeedOption)
)

func ApplyRandSeedOption(options ...RandSeedOption) (result randSeedOption) {
	for _, opt := range options {
		opt(&result)
	}
	return
}
func RandSeed(seed float64) RandSeedOption {
	return func(rso *randSeedOption) {
		rso.seed = &seed
	}
}
