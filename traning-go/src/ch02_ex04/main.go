package main

import (
	"ch02_ex04/popcount"
	"fmt"
)

func main() {
	fmt.Println(popcount.PopCount(10000))
	fmt.Println(popcount.PopCountBitShift(10000))
}
