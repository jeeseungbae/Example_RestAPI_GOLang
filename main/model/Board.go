package model

import (
	"strconv"
	"time"
)

var Boards []Board

type Board struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Writer     string    `json:"writer"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func init() {
	Boards = []Board{}

	for i := 1; i < 100; i++ {
		Boards = append(Boards, Board{
			Id:        i,
			Title:     "this is title" + strconv.Itoa(i),
			Writer:    "JSB" + strconv.Itoa(i),
			Text:      "작성 내용 입력" + strconv.Itoa(i),
			CreatedAt: time.Now(),
		})
	}
}
