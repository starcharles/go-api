package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello, world!\n")
		return
	}

	afterAll(w)
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article.... \n")
	}

	afterAll(w)
}

func ListArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article List \n")
	}

	afterAll(w)
}

func FindArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		articleId := 1
		res := fmt.Sprintf("Article No.%d\n", articleId)
		io.WriteString(w, res)
	}
	afterAll(w)
}

func PostNiceArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice... \n")
	}
	afterAll(w)
}

func PostCommnetHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Comment... \n")
	}
	afterAll(w)
}

func afterAll(w http.ResponseWriter) {
	http.Error(w, "Invalid http method", http.StatusMethodNotAllowed)
}
