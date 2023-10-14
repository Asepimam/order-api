package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Asepimam/go-micservic.git/aplications"
)

func main() {
	app := aplications.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start application", err)
	}
}
