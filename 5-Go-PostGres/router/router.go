package router

import (
	"github.com/Kdsingh333/GoLang-Project/5-Go-PostGres/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}",middleware.GetStock).Methods("GET","OPTIONS")
	router.HandleFunc("/api/stock",middleware.GetAllStock).Methods("GET")
	router.HandleFunc("/api/newStock",middleware.CreateStock).Methods("POST")
	router.HandleFunc("/api/stock/{id}",middleware.UpdateStock).Methods("PUT")
	router.HandleFunc("/api/stock/{id}",middleware.DeleteStock).Methods("DELETE")
	return router
}