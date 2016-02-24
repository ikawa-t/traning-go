package main

import (
	"bufio"
	"ch02_ex01/tmpconv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var arg string
	if len(os.Args) <= 1 {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			arg = s.Text()
			convert(arg)
			if strings.HasSuffix("\n", arg) {
				break
			}
		}
	} else {
		for _, arg := range os.Args[1:] {
			convert(arg)
		}
	}
}

func convert(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tmpconv.Fahrenheit(t)
	c := tmpconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tmpconv.FToC(f), c, tmpconv.CToF(c))
}
