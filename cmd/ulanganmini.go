package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/smkthp/ulanganmini/reader"
	Runner "github.com/smkthp/ulanganmini/runner"
	"github.com/smkthp/ulanganmini/system"
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
	chain.AddFunc(getTasks)

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

func getTasks(r Runner.Runner, ctx context.Context) error {
	r.Println("Getting Tasks...")

	tasks, err := r.Client.RunGetTasks(ctx)
	if err != nil {
		return err
	}

	if len(tasks) < 1 {
		return errors.New("no tasks yet")
	}

	displayTasks(r.Writer, tasks...)

	return nil
}

func displayTasks(w *writer.Writer, tasks ...system.Task) {
	w.Println("List of Tasks:")

	for i, task := range tasks {
		w.Println(fmt.Sprintf("%d. %s", i+1, task.Name))
	}
}
