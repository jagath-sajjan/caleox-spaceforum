package models

type Thread struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Created string `json:"created"`
	Posts   []Post `json:"posts"`
}
