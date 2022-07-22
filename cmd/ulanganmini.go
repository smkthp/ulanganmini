package main

import (
	"context"
	"fmt"

	"github.com/smkthp/ulanganmini/reader"
	"github.com/smkthp/ulanganmini/runner"
	"github.com/smkthp/ulanganmini/writer"
)

func main() {
	fmt.Println("Welcome to Ulangan Mini!")

	ctx := context.Background()
	writer := writer.NewWriter()
	reader := reader.NewReader()
	runner := runner.NewRunner(writer, reader)

	runner.Run(ctx)
}
