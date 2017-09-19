package main

import (
	"fmt"

	"github.com/Axect/Numeric/array"
	"github.com/Axect/csv"
)

func main() {
	A := make([][]float64, 2, 2)
	A[0] = array.Create(1, 1, 100)
	A[1] = array.Create(-100, 1, -1)
	B := csv.Float2Str(A)
	C := csv.Transpose(B)
	fmt.Println(C)
}
