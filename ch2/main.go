package main

import "fmt"

func main() {
	fmt.Println(AND(0, 0))
	fmt.Println(AND(1, 0))
	fmt.Println(AND(0, 1))
	fmt.Println(AND(1, 1))
}

func AND(x1, x2 float64) float64 {
	w1, w2, theta := 0.5, 0.5, 0.7
	tmp := x1*w1 + x2*w2
	if tmp <= theta {
		return 0
	}
	return 1
}
