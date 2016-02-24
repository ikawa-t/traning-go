package main

import (
	"math/rand"
	"time"
	"os"
	"./lissajous"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	Lissajous(os.Stdout)
}

