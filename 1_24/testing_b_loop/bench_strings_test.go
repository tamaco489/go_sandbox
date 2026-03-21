package testing_b_loop_test

import (
	"strings"
	"testing"
)

// BenchmarkOld は従来の b.N ループを使ったベンチマーク。
func BenchmarkOld(b *testing.B) {
	for i := 0; i < b.N; i++ { //nolint:bloop // 旧パターンの比較用
		strings.ToUpper("hello, world")
	}
}

// BenchmarkNew は Go 1.24 で追加された b.Loop() を使ったベンチマーク。
func BenchmarkNew(b *testing.B) {
	for b.Loop() {
		strings.ToUpper("hello, world")
	}
}
