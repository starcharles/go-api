package main

import (
	. "chapter2/types"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	comment1 := Comment{
		CommentId: 1,
		ArticleId: 1,
		Message:   "test comment1",
		CreatedAt: time.Now(),
	}
	comment2 := Comment{
		CommentId: 2,
		ArticleId: 1,
		Message:   "second comment",
		CreatedAt: time.Now(),
	}
	article := Article{
		Id:          1,
		Title:       "first article",
		Contents:    "This is the test article.",
		UserName:    "saki",
		NiceNum:     1,
		CommentList: []Comment{comment1, comment2},
		CreatedAt:   time.Now(),
	}

	data, err := json.Marshal(article)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Printf("%s\n", data)
}
