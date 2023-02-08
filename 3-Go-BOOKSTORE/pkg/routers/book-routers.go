package routers


import(
	"github.com/gorilla/mux"
	"github.com/Kdsingh333/GoLang-Project/3-Go-BOOKSTORE/pkg/controllers"
	
)

var RegisterBookStoreRouters = func(router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
}