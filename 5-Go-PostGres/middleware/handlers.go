package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"log"
	"net/http"
	// "os"
	"strconv"

	"github.com/Kdsingh333/GoLang-Project/5-Go-PostGres/models"
	"github.com/gorilla/mux"
	// "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")

	// }

	db, err := sql.Open("postgres", "host=localhost user=postgres password=postgres dbname=stocksdb port=5432 sslmode=disable ")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("SuccessFully connected to postgres")
	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Printf("Unabel to decode the request body, %v", err)
	}

	insertID := insertStock(stock)

	res := response{
		ID:      insertID,
		Message: "stock created successfully",
	}
	json.NewEncoder(w).Encode(res)

}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal("Error in string to int conversion")
	}
	stock, err := getStock(int64(id))

	if err != nil {
		log.Fatalf("unable to get stock, %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStock()
	if err != nil {
		log.Fatalf("unable to get all the stocks %v", err)
	}

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int %v", err)
	}

	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Unable to decode the request body %v", err)
	}

	updateRows := updateStock(int64(id), stock)

	msg := fmt.Sprintf("Stock updated successfully, total rows/records affected %v", updateRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int %v", err)

	}
	deletedRows := deleteStock(int64(id))
	msg := fmt.Sprintf("Stock deleted successfully, Total rows/records %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

// All Sub function which use above

func insertStock(stock models.Stock) int64 {

	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO stocks(name,price,company) VALUES ($1,$2,$3) RETURNING stockid`
	var id int64

	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}

func getStock(id int64) (models.Stock, error) {
	db := createConnection()
	defer db.Close()

	var stock models.Stock

	sqlStatement := `SELECT * FROM stocks WHERE stockid=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Now rows were returned")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the ro. %v", err)
	}
	return stock, err

}

func getAllStock() ([]models.Stock, error) {
	db := createConnection()
	defer db.Close()

	var stocks []models.Stock
	sqlStatement := `SELECT * FROM stocks`
	row, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	for row.Next() {
		var stock models.Stock
		err = row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatalf("Unable to scan the row %v", err)
		}
		stocks = append(stocks,stock)
	}
	return stocks,err

}

func updateStock(id int64, stock models.Stock) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement :=`UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`
	res,err := db.Exec(sqlStatement,id,stock.Name,stock.Price,stock.Company)
	if err!= nil{
		log.Fatalf("Unable to execute the query %v",err)
	}

	rowsAffected,err := res.RowsAffected()

	if err!= nil{
		log.Fatalf("Error while checking the affected rows %v",err)
		fmt.Printf("Total rows/records affected %v",rowsAffected)
	}
	return rowsAffected
}

func deleteStock(id int64) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `DELETE  FROM stocks WHERE stockid=$1`
	res,err := db.Exec(sqlStatement,id)
	if err!= nil{
		log.Fatalf("Unable to execute the query %v",err)
	}
	rowsAffected,err := res.RowsAffected()

	if err!= nil{
		log.Fatalf("Error while checking the affected rows %v",err)
		fmt.Printf("Total rows/records affected %v",rowsAffected)
	}
	return rowsAffected
}