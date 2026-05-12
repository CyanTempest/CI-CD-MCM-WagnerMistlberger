package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mrckurz/CI-CD-MCM/internal/handler"
	"github.com/mrckurz/CI-CD-MCM/internal/store"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()

	memStore := store.NewMemoryStore()
	h := handler.NewHandler(memStore)
	h.RegisterRoutes(r)

	fmt.Printf("Product Catalog API (in-memory) listening on :%s\n", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
