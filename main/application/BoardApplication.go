package application

import "main/repository"

const (
	ONE  = 1
	ZERO = 0
)

func Create(bodyData repository.Board) repository.Board {
	bodyData.Id = len(repository.Boards) + ONE
	repository.Boards = append(repository.Boards, bodyData)
	return bodyData
}

func IsBodyPresent(check repository.Board) bool {
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

func ReadAll() []repository.Board {
	return repository.Boards
}

func ReadById(sequenceNumber int) repository.Board {
	return repository.Boards[sequenceNumber-ONE]
}

func ModifyById(bodyData repository.Board, sequenceNumber int) repository.Board {
	return modify(&repository.Boards[sequenceNumber-ONE], bodyData)
}

func modify(resource *repository.Board, bodyData repository.Board) repository.Board {
	if len(bodyData.Text) != ZERO {
		resource.Text = bodyData.Text
	}
	if len(bodyData.Title) != ZERO {
		resource.Title = bodyData.Title
	}
	if len(bodyData.Writer) != ZERO {
		resource.Writer = bodyData.Writer
	}
	return bodyData
}

func DeleteById(sequenceNumber int) repository.Board {
	deleteData := repository.Boards[sequenceNumber-ONE]
	repository.Boards = deleteIndex(sequenceNumber-ONE, repository.Boards)
	return deleteData
}

func deleteIndex(index int, boards []repository.Board) []repository.Board {
	createdBoards := append(boards[:index], boards[index+ONE:]...)
	for i := 1; i <= len(createdBoards); i++ {
		createdBoards[i-ONE].Id = i
	}
	return createdBoards
}

func ValidateId(sequenceNumber int) bool {
	if len(repository.Boards) >= sequenceNumber && sequenceNumber > ZERO {
		return true
	}
	return false
}
