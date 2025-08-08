package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"catalogue-api/config"
	"catalogue-api/models"
)

// post (done)
func CreateItem(w http.ResponseWriter, r *http.Request) { //w -> used to send response back to client r -> contains information about the incoming request (url[with id], body)
	var item models.Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	//saving item into db
	result := config.Db.Create(&item)
	if result.Error != nil {
		http.Error(w, "Could not save item", http.StatusInternalServerError)
		return
	}

	//return created item as json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)

}

//get (done)

func GetItem(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "id")
	var item models.Item

	result := config.Db.First(&item, "id = ?", itemID)
	if result.Error != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	//return item as json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)

}

//put(done)

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "id") //gets item id from url
	var item models.Item

	result := config.Db.First(&item, "id = ?", itemID)
	if result.Error != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	//saving updated result
	saveResult := config.Db.Save(&item)
	if saveResult.Error != nil {
		http.Error(w, "Unable to update item", http.StatusInternalServerError)
		return
	}

	//output updated item
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)

}

// delete (done)
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "id")
	var item models.Item

	result := config.Db.First(&item, "id = ?", itemID)
	if result.Error != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	//save result of delete
	deleteResult := config.Db.Unscoped().Delete(&item)
	if deleteResult.Error != nil {
		http.Error(w, "Unable to delete item", http.StatusInternalServerError)
		return
	}

	// output success

	w.WriteHeader(http.StatusNoContent) // no content (204)

}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	var itemSlice []models.Item
	result := config.Db.Find(&itemSlice)
	if result.Error != nil {
		http.Error(w, "Could not retrieve items", http.StatusInternalServerError)
		return
	}

	// Return all items as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemSlice)
}
