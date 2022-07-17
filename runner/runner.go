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
	ctx context.Context
	client *client.Client
}

func NewRunner(ctx context.Context) *Runner {
	return &Runner{
		ctx: ctx,
		client: DefaultClient,
	}
}

func (r Runner) Run() {
	// check connection
	err := r.client.RunPing()
	if err != nil {
		log.Fatal(err)
	}
}
