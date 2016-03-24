package main

import (
	"fmt"
)

func main() {
	data := []string{"one", "two", "two", "three", "three", "three"}
	data = deleteNeighborOverlapping(data)
	fmt.Println(data)
}

// 隣接している重複を排除する
func deleteNeighborOverlapping(s []string) []string {
	out := []string{s[0]}
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			out = append(out, s[i])
		}
	}
	return out
}
