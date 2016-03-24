package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"bufio"
)

// dawnloadのAPIがわかりませんでした。。

func main() {
	os.Setenv("HTTP_PROXY", "http://p000480075:********@proxy.ricoh.co.jp:8080")
	os.Setenv("HTTPS_PROXY", "http://p000480075:********@proxy.ricoh.co.jp:8080")

	fmt.Println("映画タイトルを入力してください")
	sc := bufio.NewScanner(os.Stdin)
	var title string
	if sc.Scan() {
        title = sc.Text()
    }

	resp, err := http.Get("http://www.omdbapi.com/?t="+title+"&y=&plot=short&r=json")
	if err != nil {
		fmt.Println("http.Get err", err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll err", err)
		return
	}
	resp.Body.Close()

	fmt.Println(string(b))
}