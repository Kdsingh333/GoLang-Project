package models

import (
	"github.com/Kdsingh333/GoLang-Project/3-Go-BOOKSTORE/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(Id int64)(*Book,*gorm.DB){
	var getBook Book
	db := db.Where("ID=?",Id).Find(&getBook)
	return &getBook,db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?",ID).Delete(&book)// Using this set of code we just updated the value in delete at after we can not get it from database 
	// db.Unscoped().Where("ID=?",ID).Delete(&book)
	// Using this line of code we can drop the data from the table it permanently delete inside the database 
	return book
}