package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// x := mat.NewDense(1, 5, []float64{0.3, 1.2, -0.5, 2.0, -1.7})
	// y := StepFunction(x)
	// util.MatPrint(y)

	// x := mat.NewDense(1, 3, []float64{-1.0, 1.0, 2.0})
	// y := Sigmoid(x)
	// util.MatPrint(y)

	p := plot.New()

	p.X.Min = -6
	p.X.Max = 6
	p.Y.Min = -0.1
	p.Y.Max = 1.1

	xs := []float64{}
	for i := -5.0; i <= 5.0; i += 0.1 {
		xs = append(xs, i)
	}
	ySigmoid := StepFunction(mat.NewDense(1, len(xs), xs))
	yStepFUnction := Sigmoid(mat.NewDense(1, len(xs), xs))
	ptsSigmoid := make(plotter.XYs, len(xs))
	ptsStepFunction := make(plotter.XYs, len(xs))
	for i := range ptsSigmoid {
		ptsSigmoid[i].X = xs[i]
		ptsSigmoid[i].Y = ySigmoid.At(0, i)
		ptsStepFunction[i].X = xs[i]
		ptsStepFunction[i].Y = yStepFUnction.At(0, i)
	}
	sigmoid, err := plotter.NewLine(ptsSigmoid)
	if err != nil {
		panic(err)
	}
	stepFunction, err := plotter.NewLine(ptsStepFunction)
	if err != nil {
		panic(err)
	}
	sigmoid.LineStyle.Width = vg.Points(1)
	sigmoid.LineStyle.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	p.Add(sigmoid)
	p.Add(stepFunction)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "step_function&sigmoid.png"); err != nil {
		panic(err)
	}
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
