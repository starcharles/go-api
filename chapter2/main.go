package main

import (
	"go-api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)

	articles := router.PathPrefix("/article").Subrouter()
	articles.HandleFunc("", handlers.PostArticleHandler).Methods(http.MethodPost)
	articles.HandleFunc("/list", handlers.ListArticleHandler).Methods(http.MethodGet)
	articles.HandleFunc("/{id:[0-9]+}", handlers.FindArticleHandler).Methods(http.MethodGet)
	articles.HandleFunc("/nice", handlers.PostNiceArticleHandler).Methods(http.MethodPost)

	router.HandleFunc("/comment", handlers.PostCommnetHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
