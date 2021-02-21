package models

import "time"

//Question 问题
type Question struct {
	QuestionID uint      `json:"questionID" db:"question_id"`
	AuthorID   uint      `json:"authorID" db:"author_id"`
	CategoryID uint      `json:"categoryID" db:"category_id"`
	Status     int       `json:"status" db:"status"`
	Caption    string    `json:"caption" db:"caption"`
	Content    string    `json:"content" db:"content"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
}

//APIQuestion ...
type APIQuestion struct {
	Question
	AuthorName string `json:"authorName"`
}

//APIQuestionDetail ...
type APIQuestionDetail struct {
	Question
	AuthorName   string `json:"authorName"`
	CategoryName string `json:"categoryName"`
}

//APIQuestionList ...
type APIQuestionList struct {
	QuestionList []*APIQuestion `json:"questionList"`
	TotalCount   int            `json:"totalCount"`
}
