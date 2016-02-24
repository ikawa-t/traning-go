package main

import "testing"

// 実行コマンド
// go test -bench .

func BenchmarkInefficiencyVersion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		inefficiencyVersion()
	}
}

func BenchmarkJoinVersion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		joinVersion()
	}
}
