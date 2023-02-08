package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Kdsingh333/GoLang-Project/3-Go-BOOKSTORE/pkg/models"
	"github.com/Kdsingh333/GoLang-Project/3-Go-BOOKSTORE/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookByID(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing ")

	}
	Book := models.DeleteBook(ID)

	res, _ := json.Marshal(Book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter,r *http.Request){
     var UpdateBook = &models.Book{}
	 utils.ParseBody(r,UpdateBook);
	 vars := mux.Vars(r)
	 bookId := vars["bookId"]
	 ID,err:= strconv.ParseInt(bookId,0,0)
	 if err!= nil{
		fmt.Println("error while parsing")
	 }
	 bookDetails,db:= models.GetBookByID(ID)
	 if UpdateBook.Name != ""{
		bookDetails.Name = UpdateBook.Name
	 }
	 if UpdateBook.Author != ""{
		bookDetails.Author = UpdateBook.Author
	 }

	 if UpdateBook.Publication != ""{
		bookDetails.Publication = UpdateBook.Publication
	 }
	 db.Save(&bookDetails)
	 res,_:= json.Marshal(bookDetails)
	 w.Header().Set("Content-Type","application/json")
	 w.WriteHeader(http.StatusOK)
	 w.Write(res)
}
