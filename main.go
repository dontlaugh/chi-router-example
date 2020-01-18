package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	r.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir(filesDir)).ServeHTTP(w, r)
	}))

	port := os.Getenv("PORT")
	if port == "" {
		// unset, default to 3000
		port = "3000"
	}
	address := fmt.Sprintf(":%s", port)
	log.Println("starting server")
	http.ListenAndServe(address, r)
}
