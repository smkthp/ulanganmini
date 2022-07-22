package reader

import (
	"bytes"
	"testing"
)

func TestNewReader(t *testing.T) {
	reader := NewReader()
	if reader.reader == nil {
		t.Fatal("NewReader return writer with nil io.Writer inside it")
	}
}

func TestReadLine(t *testing.T) {
	buf := bytes.NewBuffer([]byte{'k', 'u', 'd', 'a', '\n'})
	reader := Reader {
		reader: buf,
	}
	want := "kuda"

	got := reader.ReadLine()
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
