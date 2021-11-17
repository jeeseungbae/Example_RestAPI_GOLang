package repository

import (
	"main/model"
	"time"
)

const (
	ONE  = 1
	ZERO = 0
)

func Create(bodyData model.Board) model.Board {
	bodyData.Id = len(model.Boards) + ONE
	bodyData.CreatedAt = time.Now()
	model.Boards = append(model.Boards, bodyData)
	return bodyData
}

func ReadById(sequenceNumber int) model.Board {
	return model.Boards[sequenceNumber-ONE]
}

func ReadAll() []model.Board {
	return model.Boards
}

func ModifyById(sequenceNumber int, resource model.Board) model.Board {
	model.Boards[sequenceNumber-ONE] = resource
	model.Boards[sequenceNumber-ONE].ModifiedAt = time.Now()
	return resource
}

func DeleteById(sequenceNumber int) model.Board {
	deleteData := model.Boards[sequenceNumber-ONE]
	setUpIndexId(sequenceNumber-ONE, model.Boards)
	return deleteData
}

func setUpIndexId(index int, boards []model.Board) {
	createdBoards := append(boards[:index], boards[index+ONE:]...)
	for i := ONE; i <= len(createdBoards); i++ {
		createdBoards[i-ONE].Id = i
	}
	model.Boards = createdBoards
}
