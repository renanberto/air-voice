package model

type Speech struct {
	ID           string `json:"id"`
	Content      string `json:"content"`
	Organization string `json:"organization",binding:"required"`
}
