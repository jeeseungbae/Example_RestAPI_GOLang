package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"main/application"
	"main/repository"

	"github.com/gorilla/mux"
)

const (
	ONE  = 1
	ZERO = 0
)

func Create(responseWriter http.ResponseWriter, request *http.Request) {
	var bodyData repository.Board
	json.NewDecoder(request.Body).Decode(&bodyData)

	if application.IsBodyPresent(bodyData) {
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusCreated)
		responseMessage(responseWriter, application.Create(bodyData))
		return
	}
	notCorrectDataMessage(responseWriter, request)
}

func responseMessage(responseWriter http.ResponseWriter, value interface{}) {
	json.NewEncoder(responseWriter).Encode(value)
}

func notCorrectDataMessage(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusBadRequest)
	responseMessage(responseWriter, "잘못된 값을 입력하셨습니다.")
}

func ReadAll(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseMessage(responseWriter, application.ReadAll())
}

func ReadById(responseWriter http.ResponseWriter, request *http.Request) {
	sequenceNumber, _ := strconv.Atoi(mux.Vars(request)["id"])
	if application.ValidateId(sequenceNumber) {
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)
		responseMessage(responseWriter, application.ReadById(sequenceNumber))
		return
	}
	notCorrectDataMessage(responseWriter, request)
}

func ModifyById(responseWriter http.ResponseWriter, request *http.Request) {
	sequenceNumber, _ := strconv.Atoi(mux.Vars(request)["id"])
	var bodyData repository.Board
	json.NewDecoder(request.Body).Decode(&bodyData)

	if application.ValidateId(sequenceNumber) {
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)
		responseMessage(responseWriter, application.ModifyById(bodyData, sequenceNumber))
		return
	}
	notCorrectDataMessage(responseWriter, request)
}

func DeleteById(responseWriter http.ResponseWriter, request *http.Request) {
	sequenceNumber, _ := strconv.Atoi(mux.Vars(request)["id"])

	if application.ValidateId(sequenceNumber) {
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)
		responseMessage(responseWriter, application.DeleteById(sequenceNumber))
		return
	}
	notCorrectDataMessage(responseWriter, request)
}