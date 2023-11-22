package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
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

	A := mat.NewDense(1, 3, []float64{1, 2, 3})
	B := mat.NewDense(1, 3, []float64{2, 4, 6})
	C := mat.NewDense(1, 3, []float64{0, 0, 0})
	C.Add(A, B)
	matPrint(C)
	matPrint(add(A, B))

	C.Sub(A, B)
	matPrint(C)
	matPrint(sub(A, B))

	C.MulElem(A, B)
	matPrint(C)
	matPrint(mulElem(A, B))

	C.DivElem(A, B)
	matPrint(C)
	matPrint(divElem(A, B))
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func add(a, b *mat.Dense) *mat.Dense {
	var c mat.Dense
	c.Add(a, b)
	return &c
}

func sub(a, b *mat.Dense) *mat.Dense {
	var c mat.Dense
	c.Sub(a, b)
	return &c
}

func mulElem(a, b *mat.Dense) *mat.Dense {
	var c mat.Dense
	c.MulElem(a, b)
	return &c
}

func divElem(a, b *mat.Dense) *mat.Dense {
	var c mat.Dense
	c.DivElem(a, b)
	return &c
}
