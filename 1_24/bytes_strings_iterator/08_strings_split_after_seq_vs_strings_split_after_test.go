package bytes_strings_iterator_test

import (
	"strings"
	"testing"
)

// TestStringsSplitAfterSeq_OldSplitAfter は従来の strings.SplitAfter を使った分割。
// sep を含んだまま分割する。
func TestStringsSplitAfterSeq_OldSplitAfter(t *testing.T) {
	s := "a,b,c"
	want := []string{"a,", "b,", "c"}

	got := strings.SplitAfter(s, ",")
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

// TestStringsSplitAfterSeq_New は Go 1.24 で追加された strings.SplitAfterSeq を使った分割。
// SplitAfter と同様に sep を含んだまま yield するが、[]string を事前確保しない。
func TestStringsSplitAfterSeq_New(t *testing.T) {
	s := "a,b,c"
	want := []string{"a,", "b,", "c"}

	var got []string
	for part := range strings.SplitAfterSeq(s, ",") {
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
