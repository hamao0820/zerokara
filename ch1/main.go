package main

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// a := []int{1,2,3,4,5}
	// fmt.Println(a)
	// fmt.Println(len(a))

	// a[4] = 99
	// fmt.Println(a)

	// fmt.Println(a[0:2])
	// fmt.Println(a[:3])
	// fmt.Println(a[:len(a) - 1])
	// fmt.Println(a[:len(a) - 2])

	// a := [3]int{1,2,3}
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T\n", a[0:])

	// me := map[string]int{"height": 180}
	// fmt.Println(me["height"])

	// me["weight"] = 70
	// fmt.Println(me)

	// hungry := true
	// sleepy := false

	// fmt.Println(!hungry)
	// fmt.Println(hungry && sleepy)
	// fmt.Println(hungry || sleepy)

	// for _, v := range []int{1,2,3}{
	// 	fmt.Println(v)
	// }

	// m := NewMan("David")
	// m.Hello()
	// m.Goodbye()

	// A := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
	// matPrint(A)

	// fmt.Println(A.Dims())

	// A := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
	// B := mat.NewDense(2, 2, []float64{3, 0, 0, 6})
	// C := mat.NewDense(2, 2, []float64{0, 1, 2, 0})
	// C.Add(A, B)

	// A := mat.NewDense(1, 3, []float64{1, 2, 3})
	// B := mat.NewDense(1, 3, []float64{2, 4, 6})
	// C := mat.NewDense(1, 3, []float64{0, 0, 0})
	// C.Add(A, B)
	// matPrint(C)
	// matPrint(add(A, B))

	// C.Sub(A, B)
	// matPrint(C)
	// matPrint(sub(A, B))

	// C.MulElem(A, B)
	// matPrint(C)
	// matPrint(mulElem(A, B))

	// C.DivElem(A, B)
	// matPrint(C)
	// matPrint(divElem(A, B))

	// A := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
	// matPrint(A)
	// fmt.Println(A.Dims())
	// fmt.Printf("%T\n", A.At(0, 0))

	// B := mat.NewDense(2, 2, []float64{3, 0, 0, 6})
	// matPrint(B)

	// matPrint(mul(A, B))
	// matPrint(scale(A, 2))

	// A := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
	// B := mat.NewDense(2, 1, []float64{10, 20})
	// mulElem(A, B) // panic

	// X := mat.NewDense(3, 2, []float64{51, 55, 14, 19, 0, 4})
	// matPrint(X)
	// matPrint(X.RowView(0))
	// matPrint(X.ColView(1))
	// fmt.Println(X.At(0, 1))

	// r, _ := X.Dims()
	// for i := 0; i < r; i++ {
	// 	matPrint(X.RowView(i))
	// }
	// matPrint(X)
	// fmt.Println(flatten(X))

	// matPrint(get(flatten(X), mat.NewDense(1, 3, []float64{0, 2, 4})))

	p := plot.New()
	p.Title.Text = "Sin"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.X.Min = 0
	p.X.Max = 6
	p.Y.Min = -1
	p.Y.Max = 1

	p.Add(plotter.NewFunction(math.Sin))
	// p.Add(plotter.NewFunction(math.Cos))

	p.Save(vg.Inch*5, vg.Inch*5, "sin.png")

	// x := []float64{}
	// y := []float64{}
	// for i := 0; i < 6/0.1; i++ {
	// 	x = append(x, float64(i)*0.1)
	// 	y = append(y, math.Sin(float64(i)*0.1))
	// }

	// pts := make(plotter.XYs, len(x))
	// for i := range pts {
	// 	pts[i].X = x[i]
	// 	pts[i].Y = y[i]
	// }

	// l, err := plotter.NewLine(pts)

	// if err != nil {
	// 	panic(err)
	// }
	// p.Add(l)

	// p.Save(vg.Inch*5, vg.Inch*5, "sin.png")

}

// func matPrint(X mat.Matrix) {
// 	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
// 	fmt.Printf("%v\n", fa)
// }

// func add(a, b *mat.Dense) *mat.Dense {
// 	var c mat.Dense
// 	c.Add(a, b)
// 	return &c
// }

// func sub(a, b *mat.Dense) *mat.Dense {
// 	var c mat.Dense
// 	c.Sub(a, b)
// 	return &c
// }

// func mulElem(a, b *mat.Dense) *mat.Dense {
// 	var c mat.Dense
// 	c.MulElem(a, b)
// 	return &c
// }

// func divElem(a, b *mat.Dense) *mat.Dense {
// 	var c mat.Dense
// 	c.DivElem(a, b)
// 	return &c
// }

// func mul(a, b *mat.Dense) *mat.Dense {
// 	var c mat.Dense
// 	c.Mul(a, b)
// 	return &c
// }

// func scale(a *mat.Dense, f float64) *mat.Dense {
// 	var c mat.Dense
// 	c.Scale(f, a)
// 	return &c
// }

// func flatten(a *mat.Dense) []float64 {
// 	r, c := a.Dims()
// 	fa := make([]float64, r*c)
// 	for i := 0; i < r; i++ {
// 		for j := 0; j < c; j++ {
// 			fa[i*c+j] = a.At(i, j)
// 		}
// 	}
// 	return fa
// }

// // アクセスしたい複数のindexを配列で受け取る
// func get(arr []float64, idx *mat.Dense) *mat.Dense {
// 	b := []float64{}
// 	indices := flatten(idx)
// 	for _, v := range indices {
// 		b = append(b, arr[int(v)])
// 	}

// 	return mat.NewDense(1, len(b), b)
// }
