package main

import (
	"bytes"
	"fmt"
)

func comma(str string) string {
	var buf bytes.Buffer
	for i, v := range str {
		if i != 0 && (len(str)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%s", string(v))
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("12345678"))
}
