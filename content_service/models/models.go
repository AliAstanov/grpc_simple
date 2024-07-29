package models

import (
	"time"
)

type Content struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type ContentRequest struct {
	Title string `json:"title"`
}
