package main

import (
	"context"
	"fmt"
	"time"

	"github.com/smkthp/ulanganmini/reader"
	Runner "github.com/smkthp/ulanganmini/runner"
	"github.com/smkthp/ulanganmini/writer"
)

func main() {
	fmt.Println("Welcome to Ulangan Mini!")

	ctx := context.Background()
	writer := writer.NewWriter()
	reader := reader.NewReader()
	runner := Runner.NewRunner(writer, reader)

	chain := Runner.Chain{}
	chain.AddFunc(pingServer)

	runner.SetChain(chain)
	runner.Run(ctx)
}

func pingServer(r Runner.Runner, ctx context.Context) error {
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

	err := r.Client.RunPing(ctx)
	if err != nil {
		pingOk <- false
		return err
	}

	pingOk <- true
	return nil
}
