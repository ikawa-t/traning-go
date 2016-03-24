package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	words := make(map[string]int)

	// 読み出し
	fp, err := os.Open("data/go.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		input := scanner.Text()
		strs := strings.Fields(input)
		for _, str := range strs {
			words[str]++
		}
	}
	fp.Close()

	for key, value := range words {
		fmt.Printf("%q\t%d\n", key, value)
	}
}
