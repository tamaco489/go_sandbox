package bytes_strings_iterator_test

import (
	"strings"
	"testing"
)

// TestStringsSplit_Break は strings.Split を使った早期終了の例。
// break しても Split の時点で 5行分の []string がすでに確保されている。
func TestStringsSplit_Break(t *testing.T) {
	s := "line1\nline2\nline3\nline4\nline5\n"

	// ループに入る前に全行分のスライスが確保される
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	t.Logf("ループ前に確保済みの行数: %d", len(lines))

	var processed []string
	for _, line := range lines {
		processed = append(processed, line)
		if len(processed) == 2 {
			break
		}
	}

	t.Logf("実際に処理した行: %v", processed)
	t.Logf("確保されたが処理しなかった行: %v", lines[len(processed):])
}

// TestStringsLines_Break は break によるイテレータの早期終了を確認する。
// s には 5行あるが、2行目で break するため 2行分しか yield されない。
func TestStringsLines_Break(t *testing.T) {
	s := "line1\nline2\nline3\nline4\nline5\n"

	var processed []string
	for line := range strings.Lines(s) {
		processed = append(processed, strings.TrimRight(line, "\n"))
		if len(processed) == 2 {
			break
		}
	}

	t.Logf("実際に処理した行: %v", processed)
	// strings.Lines はイテレータなので未処理の行を取り出す手段がない
	// → break した時点で line3〜line5 は yield すら行われていない
}
