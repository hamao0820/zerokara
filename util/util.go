package util

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func MatPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func Add(a, b *mat.Dense) *mat.Dense {
	var c mat.Dense
	c.Add(a, b)
	return &c
}

func Sub(a, b *mat.Dense) *mat.Dense {
	var c mat.Dense
	c.Sub(a, b)
	return &c
}

func MulElem(a, b *mat.Dense) *mat.Dense {
	var c mat.Dense
	c.MulElem(a, b)
	return &c
}

func DivElem(a, b *mat.Dense) *mat.Dense {
	var c mat.Dense
	c.DivElem(a, b)
	return &c
}

func Mul(a, b *mat.Dense) *mat.Dense {
	var c mat.Dense
	c.Mul(a, b)
	return &c
}

func Scale(a *mat.Dense, f float64) *mat.Dense {
	var c mat.Dense
	c.Scale(f, a)
	return &c
}
