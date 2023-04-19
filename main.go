package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8000", nil) // set listen port
}

func HelloServer(rw http.ResponseWriter, req *http.Request) {

	fmt.Fprint(rw, "Hello World!!")
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// func goDotEnvVariable(key string) string {

// 	// load .env file
// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	return os.Getenv(key)
// }

// func main() {
// 	// godotenv package
// 	appenv := goDotEnvVariable("APPENV")
// 	apport := goDotEnvVariable("PORT")
// 	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Update 1 dan Hello "+appenv)
// 	})
// 	fmt.Printf("Server running (port=" + apport + "), route: http://localhost:" + apport + "/hello\n")
// 	if err := http.ListenAndServe(":"+apport, nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
