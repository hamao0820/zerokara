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

	// A := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
	// B := mat.NewDense(2, 2, []float64{5, 6, 7, 8})

	// util.MatPrint(util.Mul(A, B))

	// A = mat.NewDense(2, 3, []float64{1, 2, 3, 4, 5, 6})
	// B = mat.NewDense(3, 2, []float64{1, 2, 3, 4, 5, 6})

	// util.MatPrint(util.Mul(A, B))

	// X := mat.NewDense(1, 2, []float64{1, 2})
	// fmt.Println(X.Dims())
	// W := mat.NewDense(2, 3, []float64{1, 3, 5, 2, 4, 6})
	// util.MatPrint(W)
	// fmt.Println(W.Dims())
	// Y := util.Mul(X, W)
	// util.MatPrint(Y)

	// X := mat.NewDense(1, 2, []float64{1.0, 0.5})
	// W1 := mat.NewDense(2, 3, []float64{0.1, 0.3, 0.5, 0.2, 0.4, 0.6})
	// B1 := mat.NewDense(1, 3, []float64{0.1, 0.2, 0.3})

	// fmt.Println(X.Dims())
	// fmt.Println(W1.Dims())
	// fmt.Println(B1.Dims())

	// A1 := util.Add(util.Mul(X, W1), B1)
	// Z1 := Sigmoid(A1)

	// util.MatPrint(A1)
	// util.MatPrint(Z1)

	// W2 := mat.NewDense(3, 2, []float64{0.1, 0.4, 0.2, 0.5, 0.3, 0.6})
	// B2 := mat.NewDense(1, 2, []float64{0.1, 0.2})

	// fmt.Println(Z1.Dims())
	// fmt.Println(W2.Dims())
	// fmt.Println(B2.Dims())

	// A2 := util.Add(util.Mul(Z1, W2), B2)
	// Z2 := Sigmoid(A2)

	// util.MatPrint(A2)
	// util.MatPrint(Z2)

	// W3 := mat.NewDense(2, 2, []float64{0.1, 0.3, 0.2, 0.4})
	// B3 := mat.NewDense(1, 2, []float64{0.1, 0.2})

	// fmt.Println(Z2.Dims())
	// fmt.Println(W3.Dims())
	// fmt.Println(B3.Dims())

	// A3 := util.Add(util.Mul(Z2, W3), B3)
	// Y := IdentityFunction(A3)

	// util.MatPrint(Y)

	// network := NewNetwork()
	// X := mat.NewDense(1, 2, []float64{1.0, 0.5})
	// Y := Forward(network, X)
	// util.MatPrint(Y)

	// a := mat.NewDense(1, 3, []float64{0.3, 2.9, 4.0})

	// expA := Exp(a)
	// util.MatPrint(expA)
	// sumExpA := mat.Sum(expA)
	// fmt.Println(sumExpA)
	// y := util.Scale(expA, 1/sumExpA)
	// util.MatPrint(y)
	// util.MatPrint(Softmax(a))

	a := mat.NewDense(1, 3, []float64{1010, 1000, 990})
	expX := Exp(a)
	sumExpX := mat.Sum(expX)
	util.MatPrint(util.Scale(expX, 1/sumExpX))

	c := mat.Max(a)
	expX = Exp(util.SubScalar(a, c))
	sumExpX = mat.Sum(expX)
	util.MatPrint(util.Scale(expX, 1/sumExpX))
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

func IdentityFunction(x mat.Matrix) mat.Matrix {
	return x
}

func Exp(x mat.Matrix) mat.Matrix {
	exp := func(_, _ int, v float64) float64 {
		return math.Exp(v)
	}
	var y mat.Dense
	y.Apply(exp, x)
	return &y
}

func Softmax(x mat.Matrix) mat.Matrix {
	c := mat.Max(x)
	expX := Exp(util.SubScalar(x, c))
	sumExpX := mat.Sum(expX)
	return util.Scale(expX, 1/sumExpX)
}

func NewNetwork() map[string]mat.Matrix {
	network := map[string]mat.Matrix{}
	network["W1"] = mat.NewDense(2, 3, []float64{0.1, 0.3, 0.5, 0.2, 0.4, 0.6})
	network["b1"] = mat.NewDense(1, 3, []float64{0.1, 0.2, 0.3})
	network["W2"] = mat.NewDense(3, 2, []float64{0.1, 0.4, 0.2, 0.5, 0.3, 0.6})
	network["b2"] = mat.NewDense(1, 2, []float64{0.1, 0.2})
	network["W3"] = mat.NewDense(2, 2, []float64{0.1, 0.3, 0.2, 0.4})
	network["b3"] = mat.NewDense(1, 2, []float64{0.1, 0.2})
	return network
}

func Forward(network map[string]mat.Matrix, x mat.Matrix) mat.Matrix {
	W1, W2, W3 := network["W1"], network["W2"], network["W3"]
	b1, b2, b3 := network["b1"], network["b2"], network["b3"]

	A1 := util.Add(util.Mul(x, W1), b1)
	Z1 := Sigmoid(A1)
	A2 := util.Add(util.Mul(Z1, W2), b2)
	Z2 := Sigmoid(A2)
	A3 := util.Add(util.Mul(Z2, W3), b3)
	Y := IdentityFunction(A3)

	return Y
}
