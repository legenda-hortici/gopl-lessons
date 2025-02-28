package main

import "fmt"

func main() {
	i := 0
	if itob(i) {
		fmt.Println("i is true")
	}
	fmt.Println(btoi(true))
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool {
	return i != 0
}
