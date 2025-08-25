package transport

import (
	"patricego/transport/endpoints"
	"github.com/gorilla/mux"
)

func NewRouter(taskHandler *endpoints.TaskHandler) *mux.Router {
	router := mux.NewRouter()

	
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")

	
	router.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")

	router.HandleFunc("/tasks/{resource_id}", taskHandler.GetTaskByResourceID).Methods("GET")

	
	router.HandleFunc("/tasks/{resource_id}", taskHandler.UpdateTask).Methods("PATCH")

	router.HandleFunc("/tasks/{resource_id}/status", taskHandler.TaskDelete).Methods("PATCH")


	
	

	
	return router
}


