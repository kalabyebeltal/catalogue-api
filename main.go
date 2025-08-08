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
	r.Get("/catalouge", handlers.GetAllItems)
	r.Post("/catalouge/post", handlers.CreateItem)
	r.Get("/catalouge/get{id}", handlers.GetItem)
	r.Put("/catalouge/update{id}", handlers.UpdateItem)
	r.Delete("/catalouge/delete{id}", handlers.DeleteItem)

	fmt.Println("Starting API...")
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		fmt.Printf("Error Encountered: %v", err.Error())
		return
	}

}
