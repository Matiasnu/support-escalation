package router

import (
	"go-server/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/task", middleware.GetAllEscalation).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middleware.CreateEscalation).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", middleware.EscalationComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoEscalation).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteEscalation).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTask", middleware.DeleteAllEscalation).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/app", middleware.GetAllApp).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/app", middleware.CreateApp).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/app/{id}", middleware.ModifyEscalationApp).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteApp/{id}", middleware.DeleteApp).Methods("DELETE", "OPTIONS")
	return router
}
