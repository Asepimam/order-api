package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/hallo", basicHendler)
	server := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}
	erro := server.ListenAndServe()
	if erro != nil {
		fmt.Println("failed to listen server", erro)
	}

}
func basicHendler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
