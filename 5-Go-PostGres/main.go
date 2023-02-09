package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kdsingh333/GoLang-Project/5-Go-PostGres/router"
)

func main() {
	r := router.Router()
	log.Fatal(http.ListenAndServe(":8000",r))
	fmt.Println("Starting server on the port 8000 ....... ")
}