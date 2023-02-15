package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Kdsingh333/GoLang-Project/server/39-Golang-react-todo-app/models"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection


// func LoadTheEnv() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// }

// func CreateDBInstance() {
// 	 connedtionString := os.Getenv("DB_URI")
// 	dbName := os.Getenv("DB_NAME")
// 	collName := os.Getenv("DB_COLLECTION_NAME")

// 	clientOptions := options.Client().ApplyURI(connedtionString)
// 	client, err := mongo.Connect(context.TODO(), clientOptions)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("connected to mongodb")
// 	collection := client.Database(dbName).Collection(collName)

// 	if collection != nil {
// 		fmt.Println(collection)
// 	}
// 	fmt.Println("collection instance created ")
// }
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTasks()
	json.NewEncoder(w).Encode(payload)

}
func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("task 1 is done")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.ToDoList
	fmt.Println("task 2 is done")
	json.NewDecoder(r.Body).Decode(&task)
	fmt.Println("task 3 is done")
	insertOneTask(task)
	fmt.Println("task 4 is done")
	json.NewEncoder(w).Encode(task)
	fmt.Println("task  5 is done")

}
func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func UndoTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOneTask(params["id"])

}
func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	count := deleteAllTask()

	json.NewEncoder(w).Encode(count)

}

func insertOneTask(task models.ToDoList) {
	fmt.Println("taskinside  1 is done")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	insertResult, err := collection.InsertOne(ctx, task)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted a single record", insertResult.InsertedID)
}
func getAllTasks() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}

		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}
func taskComplete(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count:", result.ModifiedCount)
}

func undoTask(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count :", result.ModifiedCount)

}

func deleteOneTask(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id":id}
	d,err:= collection.DeleteOne(context.Background(),filter)

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Deleted Document",d.DeletedCount)
}

func deleteAllTask() int64{
     d,err:= collection.DeleteMany(context.Background(),bson.D{{}}, nil)
    if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("Deleted document", d.DeletedCount)
	return d.DeletedCount;
}
