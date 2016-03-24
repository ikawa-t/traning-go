// fetchはURLにある内容を表示します

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	os.Setenv("HTTP_PROXY", "http://p000480075:********@proxy.ricoh.co.jp:8080")
	os.Setenv("HTTPS_PROXY", "http://p000480075:********@proxy.ricoh.co.jp:8080")

	for _, url := range os.Args[1:] {
		// 接頭辞がなければ追加
		pref := "http://"
		if !strings.HasPrefix(url, pref) {
			url = pref + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "Status: %v\n", resp.Status)
		fmt.Fprintf(os.Stdout, "StatusCode: %v\n", resp.StatusCode)

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
