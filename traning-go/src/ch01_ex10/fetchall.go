// fetchはURLにある内容を表示します

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strings"
)

func main() {
	os.Setenv("HTTP_PROXY", "http://p000480075:takoika18@proxy.ricoh.co.jp:8080")
	os.Setenv("HTTPS_PROXY", "http://p000480075:takoika18@proxy.ricoh.co.jp:8080")

	userFile := "data/out.txt"
	file, err := os.OpenFile(userFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer file.Close()

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// 接頭辞がなければ追加
		pref := "http://"
		if !strings.HasPrefix(url, pref) {
			url = pref + url
		}
		go fetch(url, ch) // ゴルーチンを開始
	}

	file.WriteString("--------------------------------------------------\n")
	for range os.Args[1:] {
		str := <-ch
		//str := fmt.Println(<-ch) // chチャネルから受信
		file.WriteString(str + "\n")
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // chチャネルへ送信
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 資源リークさせない
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v, url, err")
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
