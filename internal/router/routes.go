package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rashad-arbab-convictional-engineering-interview/internal/api"
	"github.com/rashad-arbab-convictional-engineering-interview/internal/datasource"
	"github.com/rashad-arbab-convictional-engineering-interview/internal/middleware"
)

type Server struct {
	Port int
}

func (server *Server) NewRouter() http.Handler {
	router := mux.NewRouter()
	router.Use(middleware.RouteLog)
	router.Use(middleware.SetHeaderJson)

	router.HandleFunc("/health", getHealth).Methods("GET")
	router.HandleFunc("/products", api.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", api.GetProduct).Methods("GET")
	router.HandleFunc("/store/inventory", api.GetInventory).Methods("GET")
	return router
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	var reqProducts []api.RequestProduct
	data, err := datasource.GetData()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(api.Error{Message: err.Error()})
		return
	}
	json.Unmarshal(data, &reqProducts)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reqProducts)
}
