package models

type Post struct {
	ID      string `json:"id"`
	Author  string `json:"author"`
	Content string `json:"content"`
	Created string `json:"created"`
}
