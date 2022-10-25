package models

import "time"

type Comment struct {
	CommentId int       `json:"comment_id"`
	ArticleId int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
