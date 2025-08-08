package main

import (
	"catalouge-api/config"
	"catalouge-api/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	config.ConnectDb()
	var r *chi.Mux = chi.NewRouter()

	//handlers
	r.Post("/catalouge", handlers.CreateItem)
	r.Get("/catalouge/{id}", handlers.GetItem)
	r.Put("/catalouge/{id}", handlers.UpdateItem)
	r.Delete("/catalouge/{id}", handlers.DeleteItem)

	fmt.Println("Starting API...")
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		fmt.Printf("Error Encountered: %v", err.Error())
		return
	}

}
