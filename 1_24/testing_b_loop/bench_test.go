package testing_b_loop_test

import (
	"os"
	"strings"
	"testing"
)

// BenchmarkOld は従来の b.N ループを使ったベンチマーク。
func BenchmarkOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.ToUpper("hello, world")
	}
}

// BenchmarkNew は Go 1.24 で追加された b.Loop() を使ったベンチマーク。
func BenchmarkNew(b *testing.B) {
	for b.Loop() {
		strings.ToUpper("hello, world")
	}
}

// BenchmarkOldWithSetup は b.N ループ + setup/cleanup の従来パターン。
// setup/cleanup をループ外に明示的に配置する必要がある。
func BenchmarkOldWithSetup(b *testing.B) {
	f, err := os.CreateTemp("", "bench_old_*.txt")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(f.Name())
	defer f.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := f.WriteString("hello\n"); err != nil {
			b.Fatal(err)
		}
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
