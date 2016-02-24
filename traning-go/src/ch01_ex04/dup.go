package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCounts := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, wordCounts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, wordCounts)
			f.Close()
		}
	}

	for word, arg := range wordCounts {
		if arg != nil {
			fmt.Printf("word : %s\n", word)
			for fileName, n := range arg {
				fmt.Printf(" %s\t%d\n", fileName, n)
			}
		}
	}
}

func countLines(f *os.File, wordCounts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		fileCounts := make(map[string]int)
		if _, ok := wordCounts[input.Text()]; ok {
			fileCounts = wordCounts[input.Text()]
		}
		var fileName = f.Name()
		fileCounts[fileName]++
		wordCounts[input.Text()] = fileCounts
	}
}
