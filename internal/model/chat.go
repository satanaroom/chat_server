package model

import "time"

type CreateChat struct {
	Usernames []string
}

type Message struct {
	Text   string
	From   string
	To     string
	SentAt time.Time
}
