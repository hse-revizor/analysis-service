package models

type ContentDto struct {
	Content string `json:"content" binding:"required"`
}
