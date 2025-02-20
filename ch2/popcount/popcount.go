package main

import (
	"fmt"
	"time"
)

// pc[i] - количество бит в i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount возвращает степень заполнения (кол-во установленных битов) значения х
func PopCount(x uint64) (int, float64) {
	start := time.Now()
	res := int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
	end := time.Since(start).Seconds()
	return res, end
}

func PopCountWithFor(x uint64) (int, float64) {
	start := time.Now()
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	end := time.Since(start).Seconds()
	return count, end
}

func main() {
	fmt.Println(PopCount(4))
	fmt.Println(PopCountWithFor(4))
}
