package main

import (
	"fmt"
	//"unicode"
)

func main() {
	s := "あ  い           うえお"
	b := []byte(s)
	fmt.Println(b)
	compressUnicode(b)
	fmt.Println(b)
}

func compressUnicode(byteSlice []byte) {
	for i := 0; i < len(byteSlice); i++ {
		if byteSlice[i] == '\t' || byteSlice[i] == '\n' || byteSlice[i] == '\v' || byteSlice[i] == '\f' || byteSlice[i] == '\r' || byteSlice[i] == ' ' {
			copy(byteSlice[i:], byteSlice[i+1:])
		}
	}
}
