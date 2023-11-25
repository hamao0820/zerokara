package main

import (
	"math"

	"github.com/hamao0820/zerokara/util"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// x := mat.NewDense(1, 5, []float64{0.3, 1.2, -0.5, 2.0, -1.7})
	// y := StepFunction(x)
	// util.MatPrint(y)

	// x := mat.NewDense(1, 3, []float64{-1.0, 1.0, 2.0})
	// y := Sigmoid(x)
	// util.MatPrint(y)

	// x := mat.NewDense(1, 5, []float64{-1.0, 1.2, 2.3, -3.3, 4.0})
	// y := ReLU(x)
	// util.MatPrint(y)

	// p := plot.New()

	// p.X.Min = -6
	// p.X.Max = 6
	// p.Y.Min = -0.1
	// p.Y.Max = 5.5

	// xs := []float64{}
	// for i := -5.0; i <= 5.0; i += 0.1 {
	// 	xs = append(xs, i)
	// }
	// y := ReLU(mat.NewDense(1, len(xs), xs))
	// pts := make(plotter.XYs, len(xs))
	// for i := range pts {
	// 	pts[i].X = xs[i]
	// 	pts[i].Y = y.At(0, i)
	// }
	// l, err := plotter.NewLine(pts)
	// if err != nil {
	// 	panic(err)
	// }
	// p.Add(l)

	// if err := p.Save(4*vg.Inch, 4*vg.Inch, "relu.png"); err != nil {
	// 	panic(err)
	// }

	A := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
	B := mat.NewDense(2, 2, []float64{5, 6, 7, 8})

	util.MatPrint(util.Mul(A, B))

	A = mat.NewDense(2, 3, []float64{1, 2, 3, 4, 5, 6})
	B = mat.NewDense(3, 2, []float64{1, 2, 3, 4, 5, 6})

	util.MatPrint(util.Mul(A, B))
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

func Sigmoid(x mat.Matrix) mat.Matrix {
	sigmoid := func(_, _ int, v float64) float64 {
		return 1 / (1 + math.Exp(-v))
	}
	var y mat.Dense
	y.Apply(sigmoid, x)
	return &y
}

func ReLU(x mat.Matrix) mat.Matrix {
	relu := func(_, _ int, v float64) float64 {
		return math.Max(0, v)
	}
	var y mat.Dense
	y.Apply(relu, x)
	return &y
}
