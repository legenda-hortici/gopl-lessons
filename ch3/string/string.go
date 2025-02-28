package main

import "fmt"

func main() {
	s := "hello world"
	fmt.Println(s, len(s))  // "hello world" 11
	fmt.Println(s[0], s[8]) // 104 114

	// получение подстроки
	fmt.Println(s[:5]) // hello
	fmt.Println(s[6:]) // world
	fmt.Println(s[:])  // hello world

	// конкатенация строк
	fmt.Println("goodbye" + s[5:]) // goodbye world

	s1 := "левая нога"
	t := s1
	s1 += ", правая нога"
	fmt.Println(s1) // левая нога, правая нога
	fmt.Println(t)  // левая нога

	// Строки неизменяемы: 
	// s[0] = 'L' - ошибка компиляции

}
