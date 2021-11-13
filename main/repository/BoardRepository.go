package repository

var Boards []Board

type Board struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Writer string `json:"writer"`
	Text   string `json:"text"`
}

func init() {
	Boards = []Board{
		{Id: 1, Title: "this is title", Writer: "JSB", Text: "작성 내용 입력"},
		{Id: 2, Title: "this is a title", Writer: "admin", Text: "작성 내용을 입력하세요."},
	}
}
