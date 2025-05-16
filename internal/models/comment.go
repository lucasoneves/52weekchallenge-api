package models

type Comment struct {
	Rating int    `json:"rating"` // 1 a 5
	Text   string `json:"text"`
}
