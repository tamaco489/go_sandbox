package bytes_strings_iterator_test

import (
	"bytes"
	"testing"
)

// TestBytesLines_OldSplit は従来の bytes.Split を使った行の走査。
func TestBytesLines_OldSplit(t *testing.T) {
	b := []byte("apple\nbanana\ncherry")
	want := [][]byte{[]byte("apple"), []byte("banana"), []byte("cherry")}

	got := bytes.Split(b, []byte("\n"))
	t.Logf("行数: %d", len(got))
	t.Logf("内容: %q", got)

	if len(got) != len(want) {
		t.Fatalf("行数: got=%d want=%d", len(got), len(want))
	}
	for i := range want {
		if !bytes.Equal(got[i], want[i]) {
			t.Errorf("[%d]: got=%q want=%q", i, got[i], want[i])
		}
	}
}

// TestBytesLines_New は Go 1.24 で追加された bytes.Lines を使った行の走査。
// bytes.Lines は改行文字を含んで yield する。
func TestBytesLines_New(t *testing.T) {
	b := []byte("apple\nbanana\ncherry")
	want := [][]byte{[]byte("apple\n"), []byte("banana\n"), []byte("cherry")}

	var got [][]byte
	for line := range bytes.Lines(b) {
		got = append(got, line)
	}
	t.Logf("行数: %d", len(got))
	t.Logf("内容: %q", got)

	if len(got) != len(want) {
		t.Fatalf("行数: got=%d want=%d", len(got), len(want))
	}
	for i := range want {
		if !bytes.Equal(got[i], want[i]) {
			t.Errorf("[%d]: got=%q want=%q", i, got[i], want[i])
		}
	}
}

// TestBytesLines_TrailingNewline は bytes.Split と bytes.Lines の末尾改行の挙動を比較する。
// bytes.Split は末尾 \n により空スライスが末尾に入るが、bytes.Lines は入らない。
func TestBytesSplit_TrailingNewline(t *testing.T) {
	b := []byte("first\nsecond\nthird\n") // 末尾の改行あり

	got := bytes.Split(b, []byte("\n"))
	t.Logf("行数: %d", len(got))
	t.Logf("内容: %q", got)
	// → ["first" "second" "third" ""] 末尾に空スライスが入る
}

// TestBytesLines_TrailingNewline は bytes.Lines の末尾改行の挙動を確認する。
// 末尾に \n があっても余分な空スライスは生成されない。
func TestBytesLines_TrailingNewline(t *testing.T) {
	b := []byte("first\nsecond\nthird\n") // 末尾の改行あり

	var got [][]byte
	for line := range bytes.Lines(b) {
		got = append(got, line)
	}
	t.Logf("行数: %d", len(got))
	t.Logf("内容: %q", got)
	// → ["first\n" "second\n" "third\n"] 末尾の空スライスなし
}
