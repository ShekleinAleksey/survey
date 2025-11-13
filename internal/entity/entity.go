package entity

import "time"

type Question struct {
	ID         int
	text       string
	created_at time.Time
}

type Answer struct {
	ID         int
	QuestionId int
	UserId     string
	Text       string
	CreatedAt  time.Time
}
