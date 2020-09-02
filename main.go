package main

import (
	"fmt"
	"net/http"
	"url_shortener/controllers"
	_ "url_shortener/models"

	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var homeHandler = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}

	hostname := os.Getenv("host")
	if hostname == "" {
		hostname = "localhost"
	}
	fmt.Fprintf(w, "<h1> URL Shortener </h1>\n <p> Enter your Shortened URL at %s:%s/XXXXXX</p>\n <h4> Created by Elijah </h4>", hostname, port)
}

func main() {

	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}

	port := os.Getenv("port")

	if port == "" {
		port = "8080"
	}
	// fmt.Println(models.GetDB())
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/{shortenedURL}", controllers.RedirectShortURL).Methods("GET")
	r.HandleFunc("/api/v1/url", controllers.CreateURLHandler).Methods("POST")
	fmt.Printf("Server started on port %s\n", port)
	http.ListenAndServe(":"+port, r)

}
