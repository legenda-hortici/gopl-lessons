package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(comma("12345"))
	fmt.Println(comma2("12345"))

	s := "abc"
	b := []byte(s)
	b = append(b, 'd')
	s2 := string(b)
	fmt.Println(s, b, s2)

	fmt.Println(comma3("abc", "cbd"))
	fmt.Println(strconv.FormatInt(12345, 2))
	y, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(y)
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// Функция добавляющая запятую при помощи буфера
func comma2(s string) string {
	var buf bytes.Buffer

	for i := 0; i <= len(s)-1; i++ {
		if i%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

// Функция проверяющая являются ли строки анаграммами
func comma3(s1, s2 string) bool {
	var flag = true
	if len(s1) == len(s2) {
		for i := 0; i <= len(s1)-1; i++ {
			fmt.Println(s1[i], s2[len(s2)-1-i])
			if s1[i] != s2[len(s2)-1-i] {
				flag = false
			}
		}
	}
	return flag
}
