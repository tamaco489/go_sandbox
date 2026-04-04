package bytes_strings_iterator_test

import (
	"strings"
	"testing"
)

// TestStringsFieldsSeq_OldFields は従来の strings.Fields を使った空白区切りの分割。
func TestStringsFieldsSeq_OldFields(t *testing.T) {
	s := "  foo   bar\tbaz  "
	want := []string{"foo", "bar", "baz"}

	got := strings.Fields(s)
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

// TestStringsFieldsSeq_New は Go 1.24 で追加された strings.FieldsSeq を使った空白区切りの分割。
// strings.Fields と同じ結果を返すが、[]string を事前確保せず 1要素ずつ yield する。
func TestStringsFieldsSeq_New(t *testing.T) {
	s := "  foo   bar\tbaz  "
	want := []string{"foo", "bar", "baz"}

	var got []string
	for field := range strings.FieldsSeq(s) {
		got = append(got, field)
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
