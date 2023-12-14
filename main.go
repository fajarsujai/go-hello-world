package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// godotenv package
	appenv := os.Getenv("APPENV")
	apport := os.Getenv("PORT")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Update 1 dan Hello "+appenv)
	})
	fmt.Printf("Server running (port=" + apport + "), route: http://localhost:" + apport + "/hello\n")
	if err := http.ListenAndServe(":"+apport, nil); err != nil {
		log.Fatal(err)
	}
}

