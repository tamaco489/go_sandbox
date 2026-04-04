package bytes_strings_iterator_test

import (
	"iter"
	"strings"
	"testing"
)

// TestIterSeq_CustomLines は strings.Lines の実際の実装を再現したカスタムイテレータ。
// IndexByte で改行位置を直接取得してスライスする実装になっている。
// 参考: https://github.com/golang/go/blob/master/src/strings/iter.go
func TestIterSeq_CustomLines(t *testing.T) {
	myLines := func(s string) iter.Seq[string] {
		return func(yield func(string) bool) {
			for len(s) > 0 {
				var line string
				if i := strings.IndexByte(s, '\n'); i >= 0 {
					line, s = s[:i+1], s[i+1:]
				} else {
					line, s = s, ""
				}
				if !yield(line) {
					return
				}
			}
		}
	}

	s := "first\nsecond\nthird"
	var got []string
	for line := range myLines(s) {
		got = append(got, line)
	}
	t.Logf("内容: %q", got)
}
