package main

import (
	"fmt"
	"math"

	"github.com/hamao0820/zerokara/matrix"
)

func main() {
	t := matrix.New1D(0, 0, 1, 0, 0, 0, 0, 0, 0, 0)
	y := matrix.New1D(0.1, 0.05, 0.6, 0, 0.05, 0.1, 0, 0.1, 0, 0)
	fmt.Println(SumSquaredError(y, t))
	fmt.Println(CrossEntropyError(y, t))

	y = matrix.New1D(0.1, 0.05, 0.1, 0, 0.05, 0.1, 0, 0.6, 0, 0)
	fmt.Println(SumSquaredError(y, t))
	fmt.Println(CrossEntropyError(y, t))
}

func SumSquaredError(y, t matrix.Matrix) float64 {
	return 0.5 * y.CopySub(t).CopyApply(func(x float64) float64 {
		return x * x
	}).Sum().At(0, 0)
}

func CrossEntropyError(y, t matrix.Matrix) float64 {
	delta := 1e-7
	return -y.CopyApply(func(x float64) float64 {
		return math.Log(x + delta)
	}).CopyMul(t).Sum().At(0, 0)
}
