package router

import (
	"github.com/gorilla/mux"
	"github.com/oricardoz/todo-service-go/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	// GET
	router.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task/{id}", middleware.GetOneTask).Methods("GET", "OPTIONS")

	// POST
	router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")

	// PUT
	router.HandleFunc("/api/task/{id}", middleware.CompleteTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/task/undo/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")

	// DELETE
	router.HandleFunc("/api/task/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/task", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")

	return router
}
