package application

import (
	"main/model"
	"main/repository"
	"strings"
)

const (
	ONE  = 1
	ZERO = 0
)

func Create(bodyData model.Board) model.Board {
	return repository.Create(bodyData)
}

func IsBodyPresent(check *model.Board) bool {
	if isEmptyData(check) {
		return IsNotDuplicateTitle(check, ZERO)
	}
	return false
}

func IsNotDuplicateTitle(data *model.Board, sequenceNumber int) bool {
	boards := repository.ReadAll()
	for i, _ := range boards {
		resourceTitle := boards[i].Title
		if strings.Contains(resourceTitle, data.Title) && i != sequenceNumber-ONE {
			return false
		}
	}
	return true
}

func isEmptyData(data *model.Board) bool {
	dataBlankRemove(data)
	if len(data.Title) == ZERO {
		return false
	}
	if len(data.Text) == ZERO {
		return false
	}
	if len(data.Writer) == ZERO {
		return false
	}
	return true
}

func dataBlankRemove(data *model.Board) {
	data.Text = strings.Trim(data.Text, " ")
	data.Title = strings.Trim(data.Title, " ")
	data.Writer = strings.Trim(data.Writer, " ")
}

func ReadAll() []model.Board {
	return repository.ReadAll()
}

func ReadById(sequenceNumber int) model.Board {
	return repository.ReadById(sequenceNumber)
}

func ModifyById(bodyData model.Board, sequenceNumber int) model.Board {
	dataBlankRemove(&bodyData)
	createdResource := modify(sequenceNumber, bodyData)
	return repository.ModifyById(sequenceNumber, createdResource)
}

func modify(sequenceNumber int, bodyData model.Board) model.Board {
	resource := repository.ReadById(sequenceNumber)
	if len(bodyData.Text) != ZERO {
		resource.Text = bodyData.Text
	}
	if len(bodyData.Title) != ZERO {
		resource.Title = bodyData.Title
	}
	if len(bodyData.Writer) != ZERO {
		resource.Writer = bodyData.Writer
	}
	return resource
}

func DeleteById(sequenceNumber int) model.Board {
	return repository.DeleteById(sequenceNumber)
}

func IsValidateId(sequenceNumber int) bool {
	boards := repository.ReadAll()
	if len(boards) >= sequenceNumber && sequenceNumber > ZERO {
		return true
	}
	return false
}
