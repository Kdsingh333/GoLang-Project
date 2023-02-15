package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kdsingh333/GoLang-Project/server/39-Golang-react-todo-app/router"
	"github.com/Kdsingh333/GoLang-Project/server/39-Golang-react-todo-app/middleware"
)

// func init() {
// 	middleware.LoadTheEnv()
// 	middleware.CreateDBInstance()

// }

func main(){
	r := router.Router() 
	fmt.Println("Starting the server on port 9000")

	log.Fatal(http.ListenAndServe(":9000",r))
}