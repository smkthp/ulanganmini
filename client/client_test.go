package client

import (
	"fmt"
	"testing"
)

func TestConcatHostPort(t *testing.T) {
	want := "http://190.0.0.1:23324"
	got := concatHostPort("http://190.0.0.1", fmt.Sprint(23324))

	if got != want {
		t.Fatalf("got: %q, want: %q", got, want)
	}
}

func BenchmarkConcatHostPort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatHostPort("http://12.112.11.1", fmt.Sprint(8837))
	}
}
