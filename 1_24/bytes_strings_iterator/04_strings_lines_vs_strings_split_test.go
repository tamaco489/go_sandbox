package bytes_strings_iterator_test

import (
	"strings"
	"testing"
)

// TestStringsLines_OldSplit は従来の strings.Split を使った行の走査。
func TestStringsLines_OldSplit(t *testing.T) {
	s := "apple\nbanana\ncherry"
	want := []string{"apple", "banana", "cherry"}

	got := strings.Split(s, "\n")
	t.Logf("行数: %d", len(got))
	t.Logf("内容: %q", got)

	if len(got) != len(want) {
		t.Fatalf("行数: got=%d want=%d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("[%d]: got=%q want=%q", i, got[i], want[i])
		}
	}
}

// TestStringsLines_New は Go 1.24 で追加された strings.Lines を使った行の走査。
// strings.Lines は改行文字を含んで yield する。
func TestStringsLines_New(t *testing.T) {
	s := "apple\nbanana\ncherry"
	want := []string{"apple\n", "banana\n", "cherry"}

	var got []string
	for line := range strings.Lines(s) {
		got = append(got, line)
	}
	t.Logf("行数: %d", len(got))
	t.Logf("内容: %q", got)

	if len(got) != len(want) {
		t.Fatalf("行数: got=%d want=%d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("[%d]: got=%q want=%q", i, got[i], want[i])
		}
	}
}
