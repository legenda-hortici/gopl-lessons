package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename2("a/b/c.go")) // c.go
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 если "/" не найден
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "/"); dot >= 0 {
		s = s[:dot]
	}
	return s
} 