package main

import (
	"fmt"
	"bytes"
	"strings"
)

func comma(str string) string {
	var buf bytes.Buffer

	// 符号があるか確認
	if strings.HasPrefix(str, "+") || strings.HasPrefix(str, "-") {
		buf.WriteByte(str[0])
		str = str[1:]
	}

	// 小数点で２分割
	strs := strings.Split(str, ".")

	// 小数点の左側
	for i, v := range strs[0] {
		if i != 0 && (len(strs[0])-i)%3 == 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%s", string(v))
	}

	// 小数点がなければここで終了
	if(len(strs) < 2) {
		return buf.String()
	}

	// 小数点を追加
	buf.WriteByte('.')

	// 小数点の右側
	for i, v := range strs[1] {
		if i != 0 && i%3 == 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%s", string(v))
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("+12345"))
	fmt.Println(comma("-1234567.89012345678"))
}

