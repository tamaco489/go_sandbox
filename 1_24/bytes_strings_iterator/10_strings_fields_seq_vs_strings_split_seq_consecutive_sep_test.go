package bytes_strings_iterator_test

import (
	"strings"
	"testing"
)

// TestStringsSplitSeq_ConsecutiveSep は strings.SplitSeq の連続セパレータの挙動を確認する。
// 連続したセパレータの間に空文字列が入る。
func TestStringsSplitSeq_ConsecutiveSep(t *testing.T) {
	s := "a,,b,,c"

	var got []string
	for part := range strings.SplitSeq(s, ",") {
		got = append(got, part)
	}
	t.Logf("要素数: %d", len(got))
	t.Logf("内容: %q", got)
	// → ["a" "" "b" "" "c"] 連続する ,, の間に空文字列が入る
}

// TestStringsFieldsSeq_ConsecutiveSep は strings.FieldsSeq の連続セパレータの挙動を確認する。
// 連続した空白はまとめて 1つの区切りとして扱われ、空文字列は生成されない。
func TestStringsFieldsSeq_ConsecutiveSep(t *testing.T) {
	s := "a  b  c"

	var got []string
	for field := range strings.FieldsSeq(s) {
		got = append(got, field)
	}
	t.Logf("要素数: %d", len(got))
	t.Logf("内容: %q", got)
	// → ["a" "b" "c"] 連続する空白をまとめて区切るため空文字列なし
}
