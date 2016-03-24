package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)	// Unicode文字の数
	var utflen [utf8.UTFMax + 1]int // UTF-8エンコーディングの長さの数
	invalid := 0 // 不正なUTF-8文字の数

	// Unicodeを文字、数字、その他に３分類する
	unicodeTypes := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // rune, nbytes, errorを返す
		if err == io.EOF { //ctl + z でぬける
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		// Unicodeを文字、数字、その他に３分類する
		if unicode.IsLetter(r) {
			unicodeTypes["Letter"]++
		}else if unicode.IsNumber(r) {
			unicodeTypes["Number"]++
		}else {
			unicodeTypes["Other"]++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i ,n)
		}
	}

	// Unicodeを文字、数字、その他に３分類する
	fmt.Printf("unicodeTypes\tcount\n")
	for c, n := range unicodeTypes {
		fmt.Printf("%q\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
