package aplications

import (
	"net/http"

	"github.com/Asepimam/go-micservic.git/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Route("/orders", LoadOrderRouter)
	return router
}

func LoadOrderRouter(r chi.Router) {
	orderHandler := &handler.Order{}

	r.Post("/", orderHandler.Create)
	r.Get("/", orderHandler.List)
	r.Get("/{id}", orderHandler.GetByID)
	r.Put("/{id}", orderHandler.UpdateByID)
	r.Delete("/{id}", orderHandler.DeleteByID)
}
