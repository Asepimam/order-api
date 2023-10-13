package main

import (
	"context"
	"fmt"

	"github.com/Asepimam/go-micservic.git/aplications"
)

func main() {
	app := aplications.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start application", err)
	}
}
