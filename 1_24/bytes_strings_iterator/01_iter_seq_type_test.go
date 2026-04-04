package bytes_strings_iterator_test

import (
	"iter"
	"testing"
)

// TestIterSeq_WhatIsIt は iter.Seq の実体を確認する。
// iter.Seq[V] は func(yield func(V) bool) という関数型のエイリアス。
// for-range に渡せる関数であれば何でも iter.Seq として扱える。
func TestIterSeq_WhatIsIt(t *testing.T) {
	// iter.Seq[string] の実体: yield を受け取り、値を1つずつ渡す関数
	var seq iter.Seq[string] = func(yield func(string) bool) {
		for _, s := range []string{"a", "b", "c"} {
			if !yield(s) {
				return // yield が false を返したら break された合図
			}
		}
	}

	var got []string
	for s := range seq {
		got = append(got, s)
	}
	t.Logf("内容: %q", got)
}
