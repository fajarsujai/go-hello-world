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
		// log.Fatal("Error loading .env file")
		log.Println("Error occurred:", err)

	}
	
	// err := doSomething()
    // if err != nil {
		//     log.Println("Error occurred:", err)
    // }
	// godotenv package
	appenv := os.Getenv("APPENV")
	apport := os.Getenv("PORT")
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	username := os.Getenv("username")
	password := os.Getenv("password")
	
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Update 1 dan Hello "+appenv)
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Nama "+name)
		fmt.Fprintf(w, "Age "+age)

	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Username "+username)
		fmt.Fprintf(w, "Password "+password)

	})

	fmt.Printf("Server running (port=" + apport + "), route: http://localhost:" + apport)
	if err := http.ListenAndServe(":"+apport, nil); err != nil {
		log.Fatal(err)
	}
}

// func doSomething() error {
//     // Contoh fungsi yang mungkin mengembalikan error
//     return nil
// }

