package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var boards []Board

const (
	ONE  = 1
	ZERO = 0
)

type Board struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Writer string `json:"writer"`
	Text   string `json:"text"`
}

func init() {
	boards = []Board{
		{Id: 1, Title: "this is title", Writer: "JSB", Text: "작성 내용 입력"},
		{Id: 2, Title: "this is a title", Writer: "admin", Text: "작성 내용을 입력하세요."},
	}
}

func readAll(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseMessage(responseWriter, boards)
}

func responseMessage(responseWriter http.ResponseWriter, value interface{}) {
	json.NewEncoder(responseWriter).Encode(value)
}

func readById(responseWriter http.ResponseWriter, request *http.Request) {
	sequenceNumber, _ := strconv.Atoi(mux.Vars(request)["id"])

	if len(boards) >= sequenceNumber && sequenceNumber > ZERO {
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)
		responseMessage(responseWriter, boards[sequenceNumber-ONE])
		return
	}
	notCorrectDataMessage(responseWriter, request)
}

func notCorrectDataMessage(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusBadRequest)
	responseMessage(responseWriter, "잘못된 값을 입력하셨습니다.")
}

func create(responseWriter http.ResponseWriter, request *http.Request) {
	var bodyData Board
	json.NewDecoder(request.Body).Decode(&bodyData)

	if isBodyPresent(bodyData) {
		bodyData.Id = len(boards) + ONE
		boards = append(boards, bodyData)

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusCreated)

		responseMessage(responseWriter, bodyData)
		return
	}
	notCorrectDataMessage(responseWriter, request)
}

func isBodyPresent(check Board) bool {
	if len(check.Title) == ZERO {
		return false
	}
	if len(check.Text) == ZERO {
		return false
	}
	if len(check.Writer) == ZERO {
		return false
	}
	return true
}

func modifyById(responseWriter http.ResponseWriter, request *http.Request) {
	sequenceNumber, _ := strconv.Atoi(mux.Vars(request)["id"])
	var bodyData Board
	json.NewDecoder(request.Body).Decode(&bodyData)

	if len(boards) > sequenceNumber {
		modify(&boards[sequenceNumber-ONE], bodyData)
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)
		responseMessage(responseWriter, boards[sequenceNumber-ONE])
		return
	}
	notCorrectDataMessage(responseWriter, request)
}

func modify(resource *Board, bodyData Board) {
	if len(bodyData.Text) != ZERO {
		resource.Text = bodyData.Text
	}
	if len(bodyData.Title) != ZERO {
		resource.Title = bodyData.Title
	}
	if len(bodyData.Writer) != ZERO {
		resource.Writer = bodyData.Writer
	}
	return
}

func deleteById(responseWriter http.ResponseWriter, request *http.Request) {
	sequenceNumber, _ := strconv.Atoi(mux.Vars(request)["id"])

	if len(boards) >= sequenceNumber {
		boards = deleteIndex(sequenceNumber - ONE)
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)
		responseMessage(responseWriter, boards[sequenceNumber-ONE])
		return
	}
	notCorrectDataMessage(responseWriter, request)
}

func deleteIndex(index int) []Board {
	createdBoards := append(boards[:index], boards[index+ONE:]...)
	for i := 1; i <= len(createdBoards); i++ {
		createdBoards[i-ONE].Id = i
	}
	return createdBoards
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/boards", create).Methods("POST")
	r.HandleFunc("/boards", readAll).Methods("GET")
	r.HandleFunc("/boards/{id:[0-9]+}", readById).Methods("GET")
	r.HandleFunc("/boards/{id:[0-9]+}", modifyById).Methods("PATCH")
	r.HandleFunc("/boards/{id:[0-9]+}", deleteById).Methods("DELETE")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
