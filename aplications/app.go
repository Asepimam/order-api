package aplications

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	Router http.Handler
	Rdb    *redis.Client
}

func New() *App {
	app := &App{
		Router: loadRoutes(),
		Rdb:    redis.NewClient(&redis.Options{}),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.Router,
	}

	erro := a.Rdb.Ping(ctx).Err()
	if erro != nil {
		return fmt.Errorf("failed to ping redis %w", erro)
	}

	defer func() {
		if erro := a.Rdb.Close(); erro != nil {
			fmt.Println("failed to close redis connection %w", erro)
		}
	}()

	chn := make(chan error, 1)

	fmt.Println("server running")
	go func() {
		erro = server.ListenAndServe()
		if erro != nil {
			chn <- fmt.Errorf("failed to listen server %w", erro)
		}
		close(chn)

	}()

	select {
	case err := <-chn:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
