package main

import (
	"go-api/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ListArticleHandler)
	http.HandleFunc("/article/1", handlers.FindArticleHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceArticleHandler)
	http.HandleFunc("/comment", handlers.PostCommnetHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
