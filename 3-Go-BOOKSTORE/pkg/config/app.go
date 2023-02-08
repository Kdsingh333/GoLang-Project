package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)
func Connect(){
	dsn := "root:@tcp(127.0.0.1:3306)/bookStore?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic(err)
	}
	db = d
	fmt.Println("connected sucessfully")
}
func GetDB() *gorm.DB{
	return db;
}