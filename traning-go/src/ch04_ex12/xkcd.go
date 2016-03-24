package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"os"
	"encoding/json"
	"bufio"
)

//const XkcdURL = "https://xkcd.com/1/info.0.json"
//const XkcdURL = "http://google.com"

type data struct {
	SafeTitle string `json:"safe_title"`
}

func main() {
	os.Setenv("HTTP_PROXY", "http://p000480075:********@proxy.ricoh.co.jp:8080")
	os.Setenv("HTTPS_PROXY", "http://p000480075:********@proxy.ricoh.co.jp:8080")

	xkcdMap := make(map[string]string)
	for i := 0; i < 10; i++ {
		title, jsonString := getXkcdJson(i)
		xkcdMap[title] = jsonString
	}

	fmt.Println("タイトル名一覧を表示します")
	for title, _  := range xkcdMap {
		fmt.Println(title)
	}

	fmt.Println("")
	fmt.Println("詳細データを表示したいコミックのタイトル名を入力してください")

	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
        s := sc.Text()
        jString, ok := xkcdMap[s]
        if !ok {
			fmt.Println("そのようなタイトルのコミックはありません")
			return
        }
        fmt.Println("コミック", s, "の詳細データは以下です")
        fmt.Println(jString)
    }
}

func getXkcdJson(comicNumber int) (title, jsonString string) {
	resp, err := http.Get("https://xkcd.com/"+strconv.Itoa(comicNumber)+"/info.0.json")
	if err != nil {
		fmt.Println("http.Get err", err)
		return "", ""
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll err", err)
		return "", ""
	}
	resp.Body.Close()

	var d data
	json.Unmarshal(b, &d)
	title = d.SafeTitle

	return title, string(b)
}
