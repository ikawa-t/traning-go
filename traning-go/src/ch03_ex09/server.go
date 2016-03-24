package main

// リクエストsample
// http://localhost:8000/surface?width=2000;height=350;color=green

import (
	"ch03_ex09/mandelbrot"
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
	http.HandleFunc("/mandelbrot", mandel)
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

func mandel(w http.ResponseWriter, r *http.Request) {
	// パラメータの取得
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	var magnification float64
	magnificationArray := r.Form["magnification"]
	if len(magnificationArray) != 0 {
		magnification, _ = strconv.ParseFloat(magnificationArray[0], 32)
	}

	// mandelbrotの表示
	mandelbrot.Mandelbrot(w, magnification)
}
