package main

import (
	"fmt"
	"math"
	"os"

	"github.com/hamao0820/zerokara/matrix"
	"github.com/sbinet/npyio"
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

	// a := mat.NewDense(1, 3, []float64{1010, 1000, 990})
	// expX := Exp(a)
	// sumExpX := mat.Sum(expX)
	// util.MatPrint(util.Scale(expX, 1/sumExpX))

	// c := mat.Max(a)
	// expX = Exp(util.SubScalar(a, c))
	// sumExpX = mat.Sum(expX)
	// util.MatPrint(util.Scale(expX, 1/sumExpX))

	// d, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	// trainDataSet, testDataSet, err := Load(d)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(trainDataSet.Count())
	// fmt.Println(testDataSet.Count())
	// img, _ := trainDataSet.Get(0)
	// fmt.Println(img.Bounds().Size())

	// file, err := os.Create("sample.png")
	// if err != nil {
	// 	log.Println("Cannot create file:", err)
	// }
	// defer file.Close()

	// png.Encode(file, img)

	// d, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	// trainDataset, testDataset, err := LoadMatrix(d)
	// if err != nil {
	// 	panic(err)
	// }
	// xTrain, tTrain := trainDataset.Images, trainDataset.Labels
	// xTest, tTest := testDataset.Images, testDataset.Labels

	// fmt.Println(xTrain.Dims())
	// fmt.Println(tTrain.Dims())
	// fmt.Println(xTest.Dims())
	// fmt.Println(tTest.Dims())

	gray := func(m matrix.Matrix) matrix.Matrix { return m.CopyDivFloat(255) }
	testSet := NewMnist(TransformData(gray))
	network, err := NewNetwork()
	if err != nil {
		panic(err)
	}

	accuracyCnt := 0
	for i := 0; i < testSet.Len(); i++ {
		img, label := testSet.Get(i)
		y := predict(network, img)
		_, c := Argmax(y)
		if c == int(label.At(0, 0)) {
			accuracyCnt++
		}
	}

	fmt.Println("Accuracy:", float64(accuracyCnt)/float64(testSet.Len()))
}

func StepFunction(x matrix.Matrix) matrix.Matrix {
	stepFunction := func(v float64) float64 {
		if v <= 0 {
			return 0
		}
		return 1
	}
	y := x.Copy()
	y.Apply(stepFunction)
	return y
}

func Sigmoid(x matrix.Matrix) matrix.Matrix {
	sigmoid := func(v float64) float64 {
		return 1 / (1 + math.Exp(-v))
	}
	return x.CopyApply(sigmoid)
}

func ReLU(x matrix.Matrix) matrix.Matrix {
	relu := func(v float64) float64 {
		return math.Max(0, v)
	}
	return x.CopyApply(relu)
}

func IdentityFunction(x matrix.Matrix) matrix.Matrix {
	return x
}

func Exp(x matrix.Matrix) matrix.Matrix {
	exp := func(v float64) float64 {
		return math.Exp(v)
	}
	return x.CopyApply(exp)
}

func Softmax(x matrix.Matrix) matrix.Matrix {
	c := x.Max().At(0, 0)
	expX := Exp(x.CopySubFloat(-c))
	sumExpX := expX.Sum().At(0, 0)
	return expX.CopyDivFloat(sumExpX)
}

func NewNetwork() (map[string]matrix.Matrix, error) {
	weights := []string{"W1", "W2", "W3", "b1", "b2", "b3"}
	network := map[string]matrix.Matrix{}
	for _, w := range weights {
		p := fmt.Sprintf("data/weights/%s.npy", w)
		f, err := os.Open(p)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		r, err := npyio.NewReader(f)
		if err != nil {
			return nil, err
		}

		shape := r.Header.Descr.Shape

		if len(shape) == 1 {
			raw := make([]float32, shape[0])
			err = r.Read(&raw)
			if err != nil {
				return nil, err
			}

			var tmp []float64
			for _, v := range raw {
				tmp = append(tmp, float64(v))
			}

			network[w] = matrix.New1D(tmp...)

			continue
		}

		raw := make([]float32, shape[0]*shape[1])
		err = r.Read(&raw)
		if err != nil {
			return nil, err
		}

		var raw64 [][]float64
		for i := 0; i < shape[0]; i++ {
			tmp := make([]float64, 0, shape[1])
			for j := 0; j < shape[1]; j++ {
				tmp = append(tmp, float64(raw[j*shape[0]+i]))
			}
			raw64 = append(raw64, tmp)
		}

		network[w] = matrix.New2D(raw64)
	}

	return network, nil
}

func Forward(network map[string]matrix.Matrix, x matrix.Matrix) matrix.Matrix {
	W1, W2, W3 := network["W1"], network["W2"], network["W3"]
	b1, b2, b3 := network["b1"], network["b2"], network["b3"]

	A1 := x.CopyMatMul(W1).CopyAdd(b1)
	Z1 := Sigmoid(A1)
	A2 := Z1.CopyMatMul(W2).CopyAdd(b2)
	Z2 := Sigmoid(A2)
	A3 := Z2.CopyMatMul(W3).CopyAdd(b3)
	Y := IdentityFunction(A3)

	return Y
}

func predict(network map[string]matrix.Matrix, x matrix.Matrix) matrix.Matrix {
	W1, W2, W3 := network["W1"], network["W2"], network["W3"]
	b1, b2, b3 := network["b1"], network["b2"], network["b3"]

	A1 := x.CopyMatMul(W1).CopyAdd(b1)
	Z1 := Sigmoid(A1)
	A2 := Z1.CopyMatMul(W2).CopyAdd(b2)
	Z2 := Sigmoid(A2)
	A3 := Z2.CopyMatMul(W3).CopyAdd(b3)
	Y := Softmax(A3)

	return Y
}

func Argmax(matrix matrix.Matrix) (int, int) {
	rows, cols := matrix.Dims()

	var maxVal float64
	var maxRow, maxCol int

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			val := matrix.At(i, j)
			if val > maxVal {
				maxVal = val
				maxRow, maxCol = i, j
			}
		}
	}

	return maxRow, maxCol
}
