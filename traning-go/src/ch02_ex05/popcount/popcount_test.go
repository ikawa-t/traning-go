package popcount

import (
	"testing"
)

// 実行コマンド
// go test -bench .

func BenchmarkPopCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCount(1000000000)
	}
}

func BenchmarkPopCountRoot(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCountRoot(1000000000)
	}
}

func BenchmarkPopCountBitShift(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCountBitShift(1000000000)
	}
}

func BenchmarkPopCountBitClear(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCountBitClear(1000000000)
	}
}

