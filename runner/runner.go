package runner

import (
	"context"
	"time"

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

type Runner struct {
	client *client.Client
	writer *writer.Writer
	reader *reader.Reader
	state  state
}

func NewRunner(writer *writer.Writer, reader *reader.Reader) *Runner {
	return &Runner{
		client: DefaultClient,
		writer: writer,
		reader: reader,
		state:  StateIdle,
	}
}

func (r *Runner) Run(ctx context.Context) {
	r.Repl(ctx)
}

func (r *Runner) Exec(ctx context.Context, str string) string {
	switch r.state {
	case StateIdle:
		r.state = StatePing
		r.Exec(ctx, str)
	case StatePing:
		err := pingServer(*r, ctx)
		if err != nil {
			r.Println("Press ENTER to try again! ")
		}
	}

	return ""
}

func (r *Runner) Repl(ctx context.Context) {
	for {
		r.Print(">>> ")
		str := r.reader.ReadLine()
		r.Println(r.Exec(ctx, str))
	}
}

func (r Runner) Print(a ...any) (n int, err error) {
	return r.writer.Print(a...)
}

func (r Runner) Println(a ...any) (n int, err error) {
	return r.writer.Println(a...)
}

func pingServer(r Runner, ctx context.Context) error {
	pingOk := make(chan bool)
	defer close(pingOk)

	r.Println("Pinging the server")

	go func() {
	p:
		for {
			select {
			case ok := <-pingOk:
				if !ok {
					r.Println("FAIL")
				}

				if ok {
					r.Println("OK")
				}

				break p
			default:
				r.Print(".")
			}
			time.Sleep(time.Millisecond * 50)
		}
	}()

	err := r.client.RunPing(ctx)
	if err != nil {
		pingOk <- false
		return err
	}

	pingOk <- true
	return nil
}
