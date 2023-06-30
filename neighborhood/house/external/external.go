package external

import (
	person "github.com/peacebaker/secretSocial/neighborhood/person/external"
)

type House struct {
	Owner    person.Person `json:"owner"`
	Posts    map[int]Post    `json:"posts"`
	Comments map[int]Comment `json:"comments"`
	Messages map[int]Message `json:"messages"`
}

// TODO: Post needs to be an interface
// text, pics, video, sound, etc can implement it 
type Post struct {
	ID            int                     `json:"id"`
	Owner         int                     `json:"owner"`
	Content       string                  `json:"content"`
	SubscriberIDs map[int]person.Person `json:"subscribers"`
}

type Comment struct {
	ID      int    `json:"id"`
	Owner   int    `json:"owner"`
	PostID  int    `json:"postid"`
	Content string `json:"content"`
}

type Message struct {
	ID      int    `json:"id"`
	Owner   int    `json:"owner"`
	ChatID  int    `json:"chatid"`
	Content string `json:"content"`
}