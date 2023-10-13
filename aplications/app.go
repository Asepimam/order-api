package aplications

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	Router http.Handler
}

func New() *App {
	app := &App{
		Router: loadRoutes(),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.Router,
	}
	erro := server.ListenAndServe()
	if erro != nil {
		fmt.Errorf("failed to listen server %w", erro)
	}
	return nil
}
