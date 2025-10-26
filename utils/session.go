package utils

import (
	"encoding/json"
	"io/ioutil"
)

type Session struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

func SaveSession(userID, username string) {
	s := Session{UserID: userID, Username: username}
	data, _ := json.Marshal(s)
	ioutil.WriteFile(".session", data, 0644)
}

func LoadSession() (*Session, error) {
	data, err := ioutil.ReadFile(".session")
	if err != nil {
		return nil, err
	}
	var s Session
	json.Unmarshal(data, &s)
	return &s, nil
}

func ClearSession() {
	ioutil.WriteFile(".session", []byte("{}"), 0644)
}
