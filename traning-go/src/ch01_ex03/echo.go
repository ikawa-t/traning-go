package main

import (
	"os"
	"strings"
)

func main() {
	//inefficiencyVersion();
	//joinVersion();
}

func inefficiencyVersion() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
}

func joinVersion() {
	strings.Join(os.Args[1:], " ")
}
