package main

import (
	"fmt"

	"github.com/hamao0820/zerokara/util"
	"gonum.org/v1/gonum/mat"
)

func main() {
	fmt.Println(AND(0, 0))
	fmt.Println(AND(1, 0))
	fmt.Println(AND(0, 1))
	fmt.Println(AND(1, 1))

	fmt.Println(NAND(0, 0))
	fmt.Println(NAND(1, 0))
	fmt.Println(NAND(0, 1))
	fmt.Println(NAND(1, 1))

	fmt.Println(OR(0, 0))
	fmt.Println(OR(1, 0))
	fmt.Println(OR(0, 1))
	fmt.Println(OR(1, 1))
}

// func AND(x1, x2 float64) float64 {
// 	w1, w2, theta := 0.5, 0.5, 0.7
// 	tmp := x1*w1 + x2*w2
// 	if tmp <= theta {
// 		return 0
// 	}
// 	return 1
// }

func AND(x1, x2 float64) float64 {
	x := mat.NewDense(2, 1, []float64{x1, x2})
	w := mat.NewDense(2, 1, []float64{0.5, 0.5})
	b := -0.7
	tmp := mat.Sum(util.MulElem(w, x)) + b
	if tmp <= 0 {
		return 0
	}
	return 1
}

func NAND(x1, x2 float64) float64 {
	x := mat.NewDense(2, 1, []float64{x1, x2})
	w := mat.NewDense(2, 1, []float64{-0.5, -0.5})
	b := 0.7
	tmp := mat.Sum(util.MulElem(w, x)) + b
	if tmp <= 0 {
		return 0
	}
	return 1
}

func OR(x1, x2 float64) float64 {
	x := mat.NewDense(2, 1, []float64{x1, x2})
	w := mat.NewDense(2, 1, []float64{0.5, 0.5})
	b := -0.2
	tmp := mat.Sum(util.MulElem(w, x)) + b
	if tmp <= 0 {
		return 0
	}
	return 1
}
