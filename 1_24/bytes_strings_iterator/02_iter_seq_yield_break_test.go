package bytes_strings_iterator_test

import (
	"iter"
	"testing"
)

// TestIterSeq_YieldAndBreak は yield と break の関係を確認する。
// break すると yield が false を返し、イテレータ側で処理を止められる。
func TestIterSeq_YieldAndBreak(t *testing.T) {
	var seq iter.Seq[string] = func(yield func(string) bool) {
		for _, s := range []string{"a", "b", "c", "d", "e"} {
			t.Logf("yield: %q", s)
			if !yield(s) {
				t.Logf("break を検知 → イテレータ終了")
				return
			}
		}
	}

	for s := range seq {
		if s == "b" {
			break // "b" の次は yield すら呼ばれない
		}
	}
}
