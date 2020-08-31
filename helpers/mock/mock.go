package mock

import (
	"fmt"
	"multi-choice/app/models"
)

type QuesService struct {
	//CreateQuestionFn func(*models.Question) (*models.Question, error)
	//UpdateQuestionFn func(*models.Question) (*models.Question, error)
	//DeleteQuestionFn func(string) error
	//GetQuestionByIDFn func(string) (*models.Question, error)
	//GetAllQuestionsFn func() ([]models.Question, error)

	CreateQuestionFn func(*models.Question) (*models.Question, error)
	UpdateQuestionFn func(*models.Question) (*models.Question, error)
	DeleteQuestionFn func(string) error
	GetByIDFn func(string) (*models.Question, error)
	GetAllQuestionsFn func() ([]*models.Question, error)
}

//type QueryInterface struct {
//	GetQuestionByIDFn func(string) (*models.Question, error)
//	GetAllQuestionsFn func() ([]models.Question, error)
//}

//create the question:
func (q *QuesService) CreateQuestion(question *models.Question) (*models.Question, error) {
	fmt.Println("ENTERED THE MOCK: ", question)
	return q.CreateQuestionFn(question)
}

func (q *QuesService) UpdateQuestion(question *models.Question) (*models.Question, error) {
	return q.UpdateQuestionFn(question)
}

func (q *QuesService) DeleteQuestion(id string) error {
	return q.DeleteQuestionFn(id)
}

func (q *QuesService) GetByID(id string) (*models.Question, error) {
	return q.GetByIDFn(id)
}

func (q *QuesService) GetAllQuestions() ([]*models.Question, error) {
	return q.GetAllQuestionsFn()
}