package routes

import (
	"go-webapp/pkg/handlers"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers all routes with the router
func RegisterRoutes(router *mux.Router) {
	// Handlers for users
	router.HandleFunc("/users", handlers.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", handlers.AddUserHandler).Methods("POST")
}
