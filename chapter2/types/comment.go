package types

import "time"

type Comment struct {
	CommentId int
	ArticleId int
	Message   string
	CreatedAt time.Time
}
