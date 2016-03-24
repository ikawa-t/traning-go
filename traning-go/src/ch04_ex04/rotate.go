package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	a = rotate(a, 2)
	fmt.Println(a)
}

func rotate(s []int, leftShift int) []int {
	if len(s) < leftShift {
		fmt.Println("input error", s, leftShift)
	}
	for i := 0; i < leftShift; i++ {
		s = append(s, s[i])
	}
	return s[leftShift:len(s)]
}
