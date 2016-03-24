package main

// リクエストsample
// http://localhost:8000/surface?width=2000;height=350;color=green

import (
	"ch03_ex04/surface"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler) // 個々のリクエストに対してhandlerが呼ばれる
	http.HandleFunc("/surface", surf)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handlerは、HTTPリクエストの情報を返します
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s &s &s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	w.Header().Set("Content-Type", "image/svg+xml")
}

func surf(w http.ResponseWriter, r *http.Request) {
	// パラメータの取得
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	var width, height float64
	var color string
	widthArray := r.Form["width"]
	if len(widthArray) != 0 {
		width, _ = strconv.ParseFloat(widthArray[0], 32)
	}
	heightArray := r.Form["height"]
	if len(heightArray) != 0 {
		height, _ = strconv.ParseFloat(heightArray[0], 32)
	}
	colorArray := r.Form["color"]
	if len(colorArray) != 0 {
		color = colorArray[0]
	}

	// surfaceの表示
	fmt.Fprintf(w, "<html>")
	fmt.Fprintf(w, "<body>")
	surface.Surface(w, width, height, color)
	fmt.Fprintf(w, "</body>")
	fmt.Fprintf(w, "</html>")
}
