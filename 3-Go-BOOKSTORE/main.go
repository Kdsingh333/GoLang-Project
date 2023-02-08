package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Kdsingh333/GoLang-Project/3-Go-BOOKSTORE/pkg/routers"
)

func main(){
	r := mux.NewRouter()
	routers.RegisterBookStoreRouters(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe(":3000",r))
}