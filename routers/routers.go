package routers

import (
	"workspace/goweb/handlers"
	"workspace/goweb/services"
	"workspace/goweb/utils"

	"github.com/gorilla/mux"
)

func GetRoute() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	router := r.PathPrefix("/lumos").Subrouter()
	router.HandleFunc("/book/{title}/{book}", handlers.GetMeTheBaby).Methods("GET")
	router.HandleFunc("/asset/asset/v1/save", utils.Chain(services.QueryReposSave, utils.Method("POST"), utils.Logging())).Methods("POST")
	router.HandleFunc("/asset/asset/v1/get-list", utils.Chain(handlers.AssetHandler, utils.Method("POST"), utils.Logging())).Methods("POST")
	router.HandleFunc("/asset/asset/v1/update", utils.Chain(services.QueryReposUpdate, utils.Method("POST"), utils.Logging())).Methods("POST")
	router.HandleFunc("/asset/asset/sqlite/v1/save", utils.Chain(services.QuerySaveSQLite, utils.Method("POST"), utils.Logging())).Methods("POST")
	router.HandleFunc("/asset/asset/sqlite/v1/get-list", utils.Chain(handlers.AssetHandlerSQLite, utils.Method("POST"), utils.Logging())).Methods("POST")
	return r
}
