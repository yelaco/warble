package entity

import "github.com/yelaco/warble/util"

type Warble struct {
	ID       string
	AuthorID string
	Body     string
}

func NewWarble(authorID, body string) Warble {
	return Warble{
		ID:       util.GenrateRandomID(),
		AuthorID: authorID,
		Body:     body,
	}
}
