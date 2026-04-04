package bytes_strings_iterator_test

import (
	"strings"
	"testing"
	"unicode"
)

// TestStringsFieldsFuncSeq_OldFieldsFunc は従来の strings.FieldsFunc を使った任意条件での分割。
// 数字を区切り文字として使用する。
func TestStringsFieldsFuncSeq_OldFieldsFunc(t *testing.T) {
	s := "abc123def456ghi"
	want := []string{"abc", "def", "ghi"}

	got := strings.FieldsFunc(s, unicode.IsDigit)
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

// TestStringsFieldsFuncSeq_New は Go 1.24 で追加された strings.FieldsFuncSeq を使った任意条件での分割。
// strings.FieldsFunc と同じ結果を返すが、[]string を事前確保せず 1要素ずつ yield する。
func TestStringsFieldsFuncSeq_New(t *testing.T) {
	s := "abc123def456ghi"
	want := []string{"abc", "def", "ghi"}

	var got []string
	for field := range strings.FieldsFuncSeq(s, unicode.IsDigit) {
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
