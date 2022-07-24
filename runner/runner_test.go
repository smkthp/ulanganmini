package runner

import (
	"bytes"
	"context"
	"testing"

	"github.com/smkthp/ulanganmini/reader"
	"github.com/smkthp/ulanganmini/writer"
)

func TestChainRun(t *testing.T) {
	type ctxKey int

	inBuf, outBuf := bytes.NewBuffer(nil), bytes.NewBuffer(nil)

	in, out := &reader.Reader{Reader: inBuf}, &writer.Writer{Writer: outBuf}
	runner := NewRunner(out, in)

	chain := Chain{}
	chainPrint := []interface{}{}

	// chain with default context
	chain.AddFunc(func(r Runner, ctx context.Context) error {
		chainPrint = append(chainPrint, ctx.Value(ctxKey(0)))
		return nil
	})

	// chain with custom context
	chain.AddFuncCtx(func(r Runner, ctx context.Context) error {
		chainPrint = append(chainPrint, ctx.Value(ctxKey(1)))
		return nil
	}, context.WithValue(context.Background(), ctxKey(1), "second func"))

	
	runner.chain = chain
	runner.Run(context.WithValue(context.Background(), ctxKey(0), "bg"))

	// first func context
	if chainPrint[0] != "bg" {
		t.Fatalf("background context got, %q, want %q", chainPrint[0], "bg")
	}

	if chainPrint[1] != "second func" {
		t.Fatalf("second func custom context got %q, want %q", chainPrint[1], "second func")
	}
}
