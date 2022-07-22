package runner

import (
	"context"
	"time"

	"github.com/smkthp/ulanganmini/client"
	"github.com/smkthp/ulanganmini/writer"
)

// api client used by runner
var DefaultClient *client.Client

func init() {
	DefaultClient = client.NewClient()
}

type Runner struct {
	client *client.Client
	writer *writer.Writer
}

func NewRunner(writer *writer.Writer) *Runner {
	return &Runner{
		client: DefaultClient,
		writer: writer,
	}
}

func (r Runner) Run(ctx context.Context) {
	pingServer(r, ctx)
	time.Sleep(time.Millisecond * 500)
}

func (r Runner) Print(a ...any) (n int, err error) {
	return r.writer.Print(a...)
}

func (r Runner) Println(a ...any) (n int, err error) {
	return r.writer.Println(a...)
}

func pingServer(r Runner, ctx context.Context) error {
	pingOk := make(chan bool)

	r.Println("Pinging the server")

	go func() {
	p:
		for {
			select {
			case ok := <- pingOk:
				if !ok {
					r.Println("FAIL")
				}

				if ok {
					r.Println("OK")
				}

				close(pingOk)
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
