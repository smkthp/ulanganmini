package writer

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Writer struct {
	mu sync.Mutex
	Writer io.Writer
}

func NewWriter() *Writer {
	return &Writer{
		Writer: os.Stdout,
	}
}

func (w *Writer) Print(a ...any) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return fmt.Fprint(w.Writer, a...)
}

func (w *Writer) Println(a ...any) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return fmt.Fprintln(w.Writer, a...)
}
