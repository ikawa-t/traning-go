package main

import (
	"bufio"
	"crypto/sha256"
	//"crypto/sha384"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
)

const (
	SHA256 int = iota
	SHA512
)

func main() {

	fmt.Println("\"SHA256\" もしくは \"SHA512\" と入力してください。")

	var shaType int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if strings.HasSuffix("SHA256", input) {
			shaType = SHA256
			break
		} else if strings.HasSuffix("SHA512", input) {
			shaType = SHA512
			break
		} else {
			fmt.Println("inputError:")
			return;
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println("文字列を入力してください")

	for scanner.Scan() {
		input := scanner.Text()
		if strings.HasSuffix("\n", input) {
			break
		}
		fmt.Println(input)
		switch shaType {
		case SHA256:
			fmt.Printf("SHA256:%x\n", sha256.Sum256([]byte(input)))
		case SHA512:
			fmt.Printf("SHA512:%x\n", sha512.Sum512([]byte(input)))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
