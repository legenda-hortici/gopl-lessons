package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "пропуск символа новой строки")
var sep = flag.String("s", " ", "разделитель")

func main() {
	p := new(int)   // р имеет тип *int и указывает на переменную типа int
	fmt.Println(*p) // разыменовываем p: 0
	*p = 2          // присваиваем p: 2
	fmt.Println(*p) // получаем новое значение p: 2

	q := new(int)
	fmt.Println(q == p) // не сравнимы т.к имеют разные адреса
}

func echo4() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
