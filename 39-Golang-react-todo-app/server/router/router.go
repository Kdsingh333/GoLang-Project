package router

import (
	
	"github.com/gorilla/mux"
	"github.com/Kdsingh333/GoLang-Project/server/39-Golang-react-todo-app/middleware"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/api/task",middleware.GetAllTasks).Methods("GET","OPTION")
	router.HandleFunc("/api/task",middleware.CreateTask).Methods("POST","OPTION")
	router.HandleFunc("/api/task{id}",middleware.TaskComplete).Methods("PUT","OPTION")
	router.HandleFunc("/api/undoTask",middleware.UndoTask).Methods("PUT","OPTION")
	router.HandleFunc("/api/deleteTask{id}",middleware.DeleteTask).Methods("DELETE","OPTION")
	router.HandleFunc("/api/deleteAllTasks",middleware.DeleteAllTask).Methods("DELETE","OPTION")
	return router
	
}