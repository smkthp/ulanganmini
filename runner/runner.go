package runner

import (
	"context"
	"log"

	"github.com/smkthp/ulanganmini/client"
)

// api client used by runner
var DefaultClient *client.Client

func init() {
	DefaultClient = client.NewClient()
}

type Runner struct {
	client *client.Client
}

func NewRunner() *Runner {
	return &Runner{
		client: DefaultClient,
	}
}

func (r Runner) Run(ctx context.Context) {
	// cancel()
	// check connection
	err := r.client.RunPing(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
