package main

// big.Floatとbig.Ratができていない

import "testing"

// 実行コマンド
// go test -bench .

func Benchmark_complex128(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mandelbrot_complex128_main()
	}
}

func Benchmark_complex64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mandelbrot_complex64_main()
	}
}

func Benchmark_bigFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mandelbrot_bigFloat_main()
	}
}
