package routes

import (
	"net/http"
	"re-write-backend/internal/handlers"
	"re-write-backend/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterPostRoutes(router *mux.Router) {
	router.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", handlers.GetPost).Methods("GET")

	// Protected routes
	router.Handle("/posts", middleware.JwtVerify(http.HandlerFunc(handlers.CreatePost))).Methods("POST")
	router.Handle("/posts/{id}", middleware.JwtVerify(http.HandlerFunc(handlers.UpdatePost))).Methods("PUT")
	router.Handle("/posts/{id}", middleware.JwtVerify(http.HandlerFunc(handlers.DeletePost))).Methods("DELETE")
}
