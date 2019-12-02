package main

import (
	"fmt"
	"github.com/toukii/analyse/pygo"
)

func main() {
	params2 := map[string]int{
		"start": 2,
		"step":  2,
		"end":   8,
	}
	args2 := [][]float64{
		[]float64{1.9, 3.9, 3.9, 3.9, 3.9, 3.9, 8.9, 9.9},
		[]float64{1.88, 3.69, 3.69, 3.69, 3.69, 3.69, 8.45, 9.36},
		[]float64{2.65, 4.59, 4.59, 4.59, 4.59, 4.59, 5.86, 7.56},
		[]float64{3.12, 4.89, 4.89, 4.89, 4.89, 4.89, 6.32, 8.52},
		[]float64{3.25, 4.56, 4.56, 4.56, 4.56, 4.56, 7.25, 9.25},
		[]float64{3.46, 4.82, 4.82, 4.82, 4.82, 4.82, 7.14, 8.89},
		[]float64{3.65, 4.15, 4.15, 4.15, 4.15, 4.15, 4.52, 7.99},
		[]float64{4.21, 4.85, 4.85, 4.85, 4.85, 4.85, 5.12, 7.65},
	}
	resp, err := pygo.GoPyFuncV2("condd", args2, params2)
	fmt.Printf("resp:%+v, err:%+v", resp, err)
}
