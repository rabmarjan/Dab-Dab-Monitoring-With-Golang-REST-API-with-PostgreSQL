package handlers

import (
	"encoding/json"
	"fmt"

	//"log"
	"net/http"
	"workspace/goweb/models"
	"workspace/goweb/services"
	"workspace/goweb/utils"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

// GetMeTheBaby calls and marshals the result as JSON
func GetMeTheBaby(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	book := vars["book"]
	log.Println("books log printed")
	fmt.Fprintf(w, "The %s is and name of book %s", title, book)
}

// AssetHandler calls `queryRepos()` and marshals the result as JSON
func AssetHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Origin", "192.168.5.156")
	//w.WriteHeader(http.StatusOK)
	utils.SetupResponse(&w, req)
	// if (*req).Method == "OPTIONS" {
	// 	return
	// }

	assets := models.Assets{}
	err := services.QueryRepos(&assets)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(assets)
	// out, err := json.Marshal(assets)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	// fmt.Fprintf(w, string(out))

}

// AssetHandlerSQLite calls `queryRepos()` and marshals the result as JSON
func AssetHandlerSQLite(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	utils.SetupResponse(&w, req)
	assets := models.Assets{}
	err := services.QueryReposSQLite(&assets)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(assets)
}
