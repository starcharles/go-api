package handlers

import (
	"encoding/json"
	"fmt"
	"go-api/models"
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
	article := models.Article1
	json, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "failed to encode to json", http.StatusInternalServerError)
	}
	w.Write(json)
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
	fmt.Println(page)

	json, err := json.Marshal([]models.Article{models.Article1, models.Article2})
	if err != nil {
		http.Error(w, "failed to encode to json", http.StatusInternalServerError)
	}
	w.Write(json)
	afterAll(w)
}

func FindArticleHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "invalid query parameter", http.StatusBadRequest)
	}
	json, err := json.Marshal(models.Article1)
	if err != nil {
		http.Error(w, "failed to encode to json", http.StatusInternalServerError)
	}
	w.Write(json)
	fmt.Println(articleId)
	// res := fmt.Sprintf("Article No.%d\n", articleId)
	// io.WriteString(w, res)
	afterAll(w)
}

func PostNiceArticleHandler(w http.ResponseWriter, req *http.Request) {
	json, err := json.Marshal(models.Article1)
	if err != nil {
		http.Error(w, "failed to encode to json", http.StatusInternalServerError)
	}
	w.Write(json)
	afterAll(w)
}

func PostCommnetHandler(w http.ResponseWriter, req *http.Request) {
	json, err := json.Marshal(models.Comment1)
	if err != nil {
		http.Error(w, "failed to encode to json", http.StatusInternalServerError)
	}
	w.Write(json)
	io.WriteString(w, "Posting Comment... \n")
	afterAll(w)
}

func afterAll(w http.ResponseWriter) {
	// Some Action
}
