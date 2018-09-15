package routers

import (
	"workspace/goweb/handlers"

	"github.com/gorilla/mux"
)

func GetRoute() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/book/{title}/{book}", handlers.GetMeTheBaby).Methods("GET")
	r.HandleFunc("/asset/asset/v1/get-list", handlers.AssetHandler).Methods("POST")
	return r
}
