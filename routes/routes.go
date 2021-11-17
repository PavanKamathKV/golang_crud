package routes

import (
	"golang_crud/services"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", services.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", services.GetTask).Methods("GET")
	router.HandleFunc("/tasks", services.CreateTask).Methods("Post")
	router.HandleFunc("/tasks/{id}", services.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", services.DeleteTask).Methods("DELETE")

	return router
}
