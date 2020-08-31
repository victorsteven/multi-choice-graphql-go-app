package question_option

import (
	"multi-choice/app/models"
)


type OptService interface {
	CreateQuestionOption(question *models.QuestionOption) (*models.QuestionOption, error)
	UpdateQuestionOption(question *models.QuestionOption) (*models.QuestionOption, error)
	DeleteQuestionOption(id string) error
	DeleteQuestionOptionByQuestionID(questionId string) error
	GetQuestionOptionByID(id string) (*models.QuestionOption, error)
	GetQuestionOptionByQuestionID(questionId string) ([]*models.QuestionOption, error)
}


