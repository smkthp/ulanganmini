package main

import (
	"context"
	"fmt"

	"github.com/smkthp/ulanganmini/runner"
)

func main() {
	fmt.Println("Welcome to Ulangan Mini!")

	ctx := context.Background()
	runner := runner.NewRunner(ctx)
	
	runner.Run()
}
