package main

import (
	"fmt"
	"net/http"

	"github.com/Kdsingh333/GoLang-Project/webCrawler-api/crawler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{vin}", crawler.CrawlForEngines)
	http.ListenAndServe(":8000", router)
	fmt.Println("We are up and running. localhost:8000/{vin}")
}

