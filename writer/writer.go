package writer

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Writer struct {
	mu sync.Mutex
	writer io.Writer
}

func NewWriter() *Writer {
	return &Writer{
		writer: os.Stdout,
	}
}

func (w *Writer) Print(a ...any) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return fmt.Fprint(w.writer, a...)
}

func (w *Writer) Println(a ...any) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return fmt.Fprintln(w.writer, a...)
}
