package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 括弧無し
	fmt.Println(strings.Join(os.Args, " "))
	// 括弧有り
	fmt.Println(strings.Join(os.Args[0:], " "))
}
