package handler

import (
	"fmt"
	h "medsearch/handler/http"
	"net/http"

	"github.com/go-chi/chi"
)

func StoreRoutes() http.Handler {
	r := chi.NewRouter()
	fmt.Println("000000000000000000000000000000000000")
	storeHandler := h.StoreHandler{}
	fmt.Println("111111111111111111111111111111111")
	r.Get("/", storeHandler.GetStore)
	// r.Post("/", bookHandler.CreateBook)
	// r.Get("/{id}", bookHandler.GetBooks)
	// r.Put("/{id}", bookHandler.UpdateBook)
	// r.Delete("/{id}", bookHandler.DeleteBook)
	return r
}
