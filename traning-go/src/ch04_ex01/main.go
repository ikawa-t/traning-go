package main

import (
	"ch04_ex01/popcount"
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	count := diffBitCount(c1, c2)
	fmt.Println("count: ", count)
}

// 異なるビット数を数えます
func diffBitCount(byteArray1 [32]byte, byteArray2 [32]byte) int {
	var count = 0
	for i := 0; i < len(byteArray1); i++ {
		count = count + popcount.BitCount(byteArray1[i], byteArray2[i])
	}
	return count
}
