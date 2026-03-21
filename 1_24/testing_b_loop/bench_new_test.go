package testing_b_loop_test

import (
	"os"
	"strings"
	"testing"
)

// BenchmarkNew は Go 1.24 で追加された b.Loop() を使ったベンチマーク。
func BenchmarkNew(b *testing.B) {
	for b.Loop() {
		strings.ToUpper("hello, world")
	}
}

// BenchmarkNewWithSetup は b.Loop() + setup/cleanup の新パターン。
// ループ外のコードは自然に setup/cleanup として機能し、計測対象が明確になる。
func BenchmarkNewWithSetup(b *testing.B) {
	f, err := os.CreateTemp("", "bench_new_*.txt")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(f.Name())
	defer f.Close()

	for b.Loop() {
		if _, err := f.WriteString("hello\n"); err != nil {
			b.Fatal(err)
		}
	}
}
