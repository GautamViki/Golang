package main

import (
	h "medsearch/handler/http"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })
    r.Mount("/med", StoreRoutes())

    http.ListenAndServe(":3000", r)
}

func StoreRoutes() http.Handler {
	r := chi.NewRouter()
	storeHandler := h.StoreHandler{}
	r.Get("/", storeHandler.GetStore)
	// r.Post("/", bookHandler.CreateBook)
	// r.Get("/{id}", bookHandler.GetBooks)
	// r.Put("/{id}", bookHandler.UpdateBook)
	// r.Delete("/{id}", bookHandler.DeleteBook)
	return r
}