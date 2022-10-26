package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-api/models"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
	afterAll(w)
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "failed to get Cotent-Length", http.StatusBadRequest)
		return
	}
	body := make([]byte, length)
	if _, err := req.Body.Read(body); !errors.Is(err, io.EOF) {
		log.Fatalln(err)
		http.Error(w, "failed to get request Body", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	fmt.Println(body)
	var reqArticle models.Article
	if err := json.Unmarshal(body, &reqArticle); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
		return
	}

	article := reqArticle
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
