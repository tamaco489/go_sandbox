package bytes_strings_iterator_test

import (
	"strings"
	"testing"
)

// TestStringsSplit_TrailingNewline は strings.Split の末尾改行の挙動を確認する。
// 末尾に \n がある場合、Split は余分な空文字列を末尾に生成する。
func TestStringsSplit_TrailingNewline(t *testing.T) {
	s := "first\nsecond\nthird\n" // 末尾の改行あり

	got := strings.Split(s, "\n")
	t.Logf("行数: %d", len(got))
	t.Logf("内容: %q", got)
	// → ["first" "second" "third" ""] 末尾に空文字列が入る
}

// TestStringsLines_TrailingNewline は strings.Lines の末尾改行の挙動を確認する。
// 末尾に \n があっても余分な空文字列は生成されない。
func TestStringsLines_TrailingNewline(t *testing.T) {
	s := "first\nsecond\nthird\n" // 末尾の改行あり

	var got []string
	for line := range strings.Lines(s) {
		got = append(got, line)
	}
	t.Logf("行数: %d", len(got))
	t.Logf("内容: %q", got)
	// → ["first\n" "second\n" "third\n"] 末尾の空文字列なし
}
