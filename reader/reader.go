package reader

import (
	"bufio"
	"io"
	"os"
	"sync"
)

type Reader struct {
	mu     sync.Mutex
	Reader io.Reader
}

func NewReader() *Reader {
	return &Reader{
		Reader: os.Stdin,
	}
}

func (r *Reader) ReadLine() (ret string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	reader := bufio.NewReader(r.Reader)
	bytes, _, err := reader.ReadLine()
	if err != nil {
		return
	}

	ret = string(bytes)
	return
}
