package bytes_strings_iterator_test

import (
	"bytes"
	"testing"
)

// TestBytesSplitSeq_OldSplit は従来の bytes.Split を使った分割。
func TestBytesSplitSeq_OldSplit(t *testing.T) {
	b := []byte("a,b,c,d")
	want := [][]byte{[]byte("a"), []byte("b"), []byte("c"), []byte("d")}

	got := bytes.Split(b, []byte(","))
	t.Logf("要素数: %d", len(got))
	t.Logf("内容: %q", got)

	if len(got) != len(want) {
		t.Fatalf("要素数: got=%d want=%d", len(got), len(want))
	}
	for i := range want {
		if !bytes.Equal(got[i], want[i]) {
			t.Errorf("[%d]: got=%q want=%q", i, got[i], want[i])
		}
	}
}

// TestBytesSplitSeq_New は Go 1.24 で追加された bytes.SplitSeq を使った分割。
// bytes.Split と同じ結果を返すが、[][]byte を事前確保せず 1要素ずつ yield する。
func TestBytesSplitSeq_New(t *testing.T) {
	b := []byte("a,b,c,d")
	want := [][]byte{[]byte("a"), []byte("b"), []byte("c"), []byte("d")}

	var got [][]byte
	for part := range bytes.SplitSeq(b, []byte(",")) {
		got = append(got, part)
	}
	t.Logf("要素数: %d", len(got))
	t.Logf("内容: %q", got)

	if len(got) != len(want) {
		t.Fatalf("要素数: got=%d want=%d", len(got), len(want))
	}
	for i := range want {
		if !bytes.Equal(got[i], want[i]) {
			t.Errorf("[%d]: got=%q want=%q", i, got[i], want[i])
		}
	}
}
