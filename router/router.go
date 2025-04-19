package router

import (
	"github.com/gorilla/mux"
	"github.com/oricardoz/todo-service-go/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("api/todos", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("api/todos/{id}", middleware.CompleteTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("api/undoTodos/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("api/todos", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("api/todos/{id}", middleware.DeleteTask).Methods("POST", "OPTIONS")
	router.HandleFunc("api/deleteAll/task", middleware.DeleteAllTask).Methods("DELET", "OPTIONS")
	return router
}
