package main

import (
	"ch02_ex03/popcount"
	"fmt"
)

func main() {
	fmt.Println(popcount.PopCount(0))
	fmt.Println(popcount.PopCount(1))
	fmt.Println(popcount.PopCount(2))
	fmt.Println(popcount.PopCount(3))
	fmt.Println(popcount.PopCount(4))
	fmt.Println(popcount.PopCount(5))
	fmt.Println(popcount.PopCount(6))
	fmt.Println(popcount.PopCount(7))
	fmt.Println(popcount.PopCount(8))
	fmt.Println(popcount.PopCount(9))
	fmt.Println(popcount.PopCount(10))
	fmt.Println(popcount.PopCount(1000000000))

	fmt.Println(popcount.PopCountRoot(0))
	fmt.Println(popcount.PopCountRoot(1))
	fmt.Println(popcount.PopCountRoot(2))
	fmt.Println(popcount.PopCountRoot(3))
	fmt.Println(popcount.PopCountRoot(4))
	fmt.Println(popcount.PopCountRoot(5))
	fmt.Println(popcount.PopCountRoot(6))
	fmt.Println(popcount.PopCountRoot(7))
	fmt.Println(popcount.PopCountRoot(8))
	fmt.Println(popcount.PopCountRoot(9))
	fmt.Println(popcount.PopCountRoot(10))
	fmt.Println(popcount.PopCountRoot(1000000000))
}
