package main

import (
	"github.com/hamao0820/zerokara/util"
	"gonum.org/v1/gonum/mat"
)

func main() {
	x := mat.NewDense(1, 5, []float64{0.3, 1.2, -0.5, 2.0, -1.7})
	y := StepFunction(x)
	util.MatPrint(y)
}

func StepFunction(x mat.Matrix) mat.Matrix {
	stepFunction := func(_, _ int, v float64) float64 {
		if v <= 0 {
			return 0
		}
		return 1
	}
	var y mat.Dense
	y.Apply(stepFunction, x)
	return &y
}
