package main

import (
	"catalogue-api/config"
	"catalogue-api/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	config.ConnectDb()
	var r *chi.Mux = chi.NewRouter()

	//handlers
	r.Get("/catalogue", handlers.GetAllItems)
	r.Post("/catalogue/post", handlers.CreateItem)
	r.Get("/catalogue/get{id}", handlers.GetItem)
	r.Put("/catalogue/update{id}", handlers.UpdateItem)
	r.Delete("/catalogue/delete{id}", handlers.DeleteItem)

	fmt.Println("Starting API...")
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		fmt.Printf("Error Encountered: %v", err.Error())
		return
	}

}
