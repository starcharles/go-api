package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
	afterAll(w)
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article.... \n")
	afterAll(w)
}

func ListArticleHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if queryMap.Has("page") {
		var err error

		page, err = strconv.Atoi(queryMap.Get("page"))

		if err != nil {
			http.Error(w, "invalid query params", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	io.WriteString(w, fmt.Sprintf("Article List (page = %d)\n", page))
	afterAll(w)
}

func FindArticleHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "invalid query parameter", http.StatusBadRequest)
	}
	res := fmt.Sprintf("Article No.%d\n", articleId)
	io.WriteString(w, res)
	afterAll(w)
}

func PostNiceArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice... \n")
	afterAll(w)
}

func PostCommnetHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment... \n")
	afterAll(w)
}

func afterAll(w http.ResponseWriter) {
	// Some Action
}
