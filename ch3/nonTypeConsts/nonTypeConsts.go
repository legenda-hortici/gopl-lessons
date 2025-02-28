package main

import (
	"fmt"
	"math"
)

func main() {
	const Pi64 = math.Pi

	var x float32 = float32(Pi64)
	var y float64 = float64(Pi64)
	var z complex128 = complex128(Pi64)

	fmt.Println(x, y, z)

	var f1 float64 = 3 + 0i // float64 - типизированная комплексное
	f1 = 2 // float64 - нетипизированное целое
	f1 = 1e123 // float64 - нетипизированное действительное
	f1 = 'a' // float64 - нетипизированная руна
	fmt.Println(f1)

	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)     // "100"; тип float64
	fmt.Println(5 / 9 * (f - 32))     // "0"; 5 / 9 нетипизированное целое (0)
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0 / 9.0 нетипизированное вещественное
}
