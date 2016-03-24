package main

import (
	"fmt"
)

func main() {
	str := "あいうえお"
	byteSlice := []byte(str)
	fmt.Println(byteSlice)
	byteSlice = reverse(byteSlice)
	fmt.Println(byteSlice)
}

func reverse(byteSlice []byte) []byte{
	ru := []rune(string(byteSlice))
	for i := 0; i < len(ru)/2; i++ {
		ru[i], ru[len(ru)-1-i] = ru[len(ru)-1-i], ru[i]
	}
	return []byte(string(ru))

}

