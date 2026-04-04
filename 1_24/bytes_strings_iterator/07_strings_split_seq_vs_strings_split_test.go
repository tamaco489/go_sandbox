package bytes_strings_iterator_test

import (
	"strings"
	"testing"
)

// TestStringsSplitSeq_OldSplit は従来の strings.Split を使った文字列の分割。
func TestStringsSplitSeq_OldSplit(t *testing.T) {
	s := "a,b,c,d"
	want := []string{"a", "b", "c", "d"}

	got := strings.Split(s, ",")
	t.Logf("要素数: %d", len(got))
	t.Logf("内容: %q", got)

	if len(got) != len(want) {
		t.Fatalf("要素数: got=%d want=%d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("[%d]: got=%q want=%q", i, got[i], want[i])
		}
	}
}

// TestStringsSplitSeq_New は Go 1.24 で追加された strings.SplitSeq を使った文字列の分割。
// strings.Split と異なり、[]string を事前確保せず 1要素ずつ yield する。
func TestStringsSplitSeq_New(t *testing.T) {
	s := "a,b,c,d"
	want := []string{"a", "b", "c", "d"}

	var got []string
	for part := range strings.SplitSeq(s, ",") {
		got = append(got, part)
	}
	t.Logf("要素数: %d", len(got))
	t.Logf("内容: %q", got)

	if len(got) != len(want) {
		t.Fatalf("要素数: got=%d want=%d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("[%d]: got=%q want=%q", i, got[i], want[i])
		}
	}
}
