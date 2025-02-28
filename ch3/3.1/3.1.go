package main

import "fmt"

func test3() {
	// logicOperators()
	transformUint8()
}

func logicOperators() {
	x, y := 5, 3
	fmt.Println(x & y)    // Побитовое И (0101 & 0011 = 0001 -> 1).
	fmt.Println(x | y)    // Побитовое ИЛИ (0101 | 0011 = 0111 -> 7)
	fmt.Println(x ^ y)    // Побитовое исключающее ИЛИ (0101 ^ 0011 = 0110 -> 6)
	fmt.Println(x << 1)   // Сдвиг влево (0101 -> 1010 -> 10)
	fmt.Println(x >> 1)   // Сдвиг вправо (0101 -> 0010 -> 2)
	fmt.Println(^uint(0)) // 18446744073709551615
}

func transformUint8() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)    // "00100010", множество {1, 5}
	fmt.Printf("%08b\n", y)    // "00100010", множество {1, 2}
	fmt.Printf("%08b\n", x&y)  // "00000010", пересечение {1}
	fmt.Printf("%08b\n", x|y)  // "00000010", объединение {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", симметричная разность {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", разность {5}

	for i := uint(0); i < 8; i++ {
		if x & (1 << i) != 0 { // Проверка принадлежности множеству
			fmt.Println(i) // "1", "5"
		}
	}
	fmt.Printf("%08b\n", x << 1) // "01000100", множество {2, 6}
	fmt.Printf("%08b\n", x >> 1) // "00010001", множество {0, 4}
	fmt.Println(0Xdeadbeef)
}
