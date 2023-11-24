package main

import (
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// x := mat.NewDense(1, 5, []float64{0.3, 1.2, -0.5, 2.0, -1.7})
	// y := StepFunction(x)
	// util.MatPrint(y)

	p := plot.New()
	p.Title.Text = "Step Function"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	p.X.Min = -6
	p.X.Max = 6
	p.Y.Min = -0.1
	p.Y.Max = 1.1

	xs := []float64{}
	for i := -5.0; i <= 5.0; i += 0.1 {
		xs = append(xs, i)
	}
	y := StepFunction(mat.NewDense(1, len(xs), xs))
	pts := make(plotter.XYs, len(xs))
	for i := range pts {
		pts[i].X = xs[i]
		pts[i].Y = y.At(0, i)
	}
	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "step_function.png"); err != nil {
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
