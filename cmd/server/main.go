package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"re-write-backend/internal/db"
	"re-write-backend/internal/routes"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	client := db.GetMongoClient()
	defer func() {
		if err := client.Disconnect(nil); err != nil {
			log.Println("Error disconnecting MongoDB:", err)
		}
	}()

	router := mux.NewRouter()

	// Register auth and post routes
	routes.RegisterAuthRoutes(router)
	routes.RegisterPostRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
