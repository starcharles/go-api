package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article.... \n")
}

func ListArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List \n")
}

func FindArticleHandler(w http.ResponseWriter, req *http.Request) {
	articleId := 1
	res := fmt.Sprintf("Article No.%d\n", articleId)
	io.WriteString(w, res)
}

func PostNiceArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice... \n")
}

func PostCommnetHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment... \n")
}
