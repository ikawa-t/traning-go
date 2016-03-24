package main

import (
	"ch01_ex12/lissajous"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler) // 個々のリクエストに対してhandlerが呼ばれる
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", liss)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handlerは、HTTPリクエストの情報を返します
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s &s &s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	//fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	fmt.Fprintf(w, "Lissajou!!!!!!!!s")
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	var cycles float64
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		if k == "cycles" {
			fmt.Fprintf(w, "cycles = %q\n", v)
		}
	}
	lissajous.Lissajous(w,cycles)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func liss(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "lissajous\n")
	//lissajous.Lissajous(w)
}
