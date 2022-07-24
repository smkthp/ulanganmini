package writer

import (
	"bytes"
	"testing"
)

func TestNewWriter(t *testing.T) {
	writer := NewWriter()
	if writer.Writer == nil {
		t.Fatal("NewWriter return writer with nil io.Writer inside it")
	}
}

func TestPrint(t *testing.T) {
	want := "test"
	buf := bytes.NewBuffer(nil)
	writer := &Writer{Writer: buf}
	writer.Print(want)

	got := buf.String()

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestPrintln(t *testing.T) {
	want := "test\n"
	buf := bytes.NewBuffer(nil)
	writer := &Writer{Writer: buf}
	writer.Println("test")

	got := buf.String()

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
} 
