package runner

import (
	"context"
	"errors"
	"os"

	"github.com/smkthp/ulanganmini/client"
	"github.com/smkthp/ulanganmini/reader"
	"github.com/smkthp/ulanganmini/writer"
)

// api client used by runner
var DefaultClient *client.Client

func init() {
	DefaultClient = client.NewClient()
}

type state int

const (
	StateBoot state = iota
	StateIdle
	StatePing
	StateGetTasks
)

type RunnerFunc func(Runner, context.Context) error

type Chain struct {
	fns []RunnerFunc
	// indicates the func index in fns that will be processed
	index int
}

// add RunnerFunc to the chain without adding context.
// the f will be called with context from Run(`ctx`)
func (c *Chain) AddFunc(f RunnerFunc) {
	c.fns = append(c.fns, f)
}

func (c *Chain) AddFuncCtx(f RunnerFunc, ctx context.Context) {
	callback := func(r Runner, _ context.Context) error {
		return f(r, ctx)
	}

	c.fns = append(c.fns, callback)
}

// Checking if the fns[index] exist. If does not exist will return error
func (c *Chain) Next() error {
	if len(c.fns) > c.index {
		return nil
	}

	return errors.New("index out of bounds")
}

// Takes function from chain based on the internal index.
// Make sure call Next() first to verify if next function available.
// Returns RunnerFunc and its index
func (c *Chain) GetFunc() (RunnerFunc, int) {
	return c.fns[c.index], c.index
}

// Will set index to index+1
func (c *Chain) Done() {
	c.index = c.index + 1
}

type Runner struct {
	Client *client.Client
	Writer *writer.Writer
	Reader *reader.Reader
	state  state
	chain  Chain
}

func NewRunner(writer *writer.Writer, reader *reader.Reader) *Runner {
	return &Runner{
		Client: DefaultClient,
		Writer: writer,
		Reader: reader,
		state:  StateIdle,
	}
}

func (r *Runner) Run(ctx context.Context) {
	for {
		if err := r.chain.Next(); err != nil {
			os.Exit(0)
		}

		fn, _ := r.chain.GetFunc()
		
		if err := fn(*r, ctx); err != nil {
			r.Println("Hit Enter to try again!")
			r.Prompt("")
			continue
		}

		r.chain.Done()
	}

}

func (r *Runner) SetChain(chain Chain) {
	r.chain = chain
}

func (r *Runner) Prompt(prefix string) string {
	r.Print(prefix)
	r.Print(">>> ")
	str := r.Reader.ReadLine()
	return str
}

func (r Runner) Print(a ...any) (n int, err error) {
	return r.Writer.Print(a...)
}

func (r Runner) Println(a ...any) (n int, err error) {
	return r.Writer.Println(a...)
}
