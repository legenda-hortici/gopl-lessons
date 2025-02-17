// echo1 выводит элкменты командной строки
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	secs := time.Since(start).Seconds()
	fmt.Println(s, "|", secs)
}

func echo2() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	secs := time.Since(start).Seconds()
	fmt.Println(s, "|", secs)
}

func echo3() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "), "|", time.Since(start).Seconds())
}

func echo() {
	echo1()
	echo2()
	echo3()
	// benchmark(nil)
}
