package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"workspace/goweb/models"
	"workspace/goweb/services"

	"github.com/gorilla/mux"
)

func GetMeTheBaby(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	book := vars["book"]
	log.Println("books log printed")
	fmt.Fprintf(w, "The %s is and name of book %s", title, book)
}

// AssetHandler calls `QueryRepos()` and marshals the result as JSON
func AssetHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)

	assets := models.Assets{}
	err := services.QueryRepos(&assets)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(assets)

// 	out, err := json.Marshal(assets)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	fmt.Fprintf(w, string(out))
}
