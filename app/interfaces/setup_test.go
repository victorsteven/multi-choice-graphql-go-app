package interfaces_test

import (
	"multi-choice/app/models"
)


//We need to mock the domain layer, so we can achieve unit test in the interfaces layer:

//Question Domain Mocking
var (
	CreateQuestionFn func(*models.Question) (*models.Question, error)
	UpdateQuestionFn func(*models.Question) (*models.Question, error)
	DeleteQuestionFn func(string) error
	GetQuestionByIDFn func(string) (*models.Question, error)
	GetAllQuestionFn func() ([]*models.Question, error)
)

func (q *fakeQuestionService) CreateQuestion(question *models.Question) (*models.Question, error) {
	return CreateQuestionFn(question)
}

func (q *fakeQuestionService) UpdateQuestion(question *models.Question) (*models.Question, error) {
	return UpdateQuestionFn(question)
}

func (q *fakeQuestionService) DeleteQuestion(id string) error {
	return DeleteQuestionFn(id)
}

func (q *fakeQuestionService) GetQuestionByID(id string) (*models.Question, error) {
	return GetQuestionByIDFn(id)
}

func (q *fakeQuestionService) GetAllQuestion() ([]*models.Question, error) {
	return GetAllQuestionFn()
}

////////////////////////////////////////
//QuestionOption Domain Mocking
var (
	CreateQuestionOptionFn func(option *models.QuestionOption) (*models.QuestionOption, error)
	UpdateQuestionOptionFn func(option *models.QuestionOption) (*models.QuestionOption, error)
	DeleteQuestionOptionFn func(string) error
	DeleteByQuestionIDFn func(questionId string) error
	GetQuestionOptionByIDFn func(string) (*models.QuestionOption, error)
	GetByQuestionID func(questionId string) ([]*models.QuestionOption, error)
)

func (o *fakeQuestionOptionService) CreateQuestionOption(option *models.QuestionOption) (*models.QuestionOption, error) {
	return CreateQuestionOptionFn(option)
}

func (o *fakeQuestionOptionService) UpdateQuestionOption(option *models.QuestionOption) (*models.QuestionOption, error) {
	return UpdateQuestionOptionFn(option)
}

func (o *fakeQuestionOptionService) GetQuestionOptionByID(id string) (*models.QuestionOption, error) {
	return GetQuestionOptionByIDFn(id)
}

func (o *fakeQuestionOptionService) DeleteQuestionOption(id string) error {
	return DeleteQuestionOptionFn(id)
}

func (o *fakeQuestionOptionService) DeleteByQuestionID(questionId string) error {
	return DeleteByQuestionIDFn(questionId)
}

func (o *fakeQuestionOptionService) GetByQuestionID(questionId string) ([]*models.QuestionOption, error) {
	return GetByQuestionID(questionId)
}


////////////////////////////////////////
//QuestionOption Domain Mocking
var (
	CreateAnswerFn func(answer *models.Answer) (*models.Answer, error)
	UpdateAnswerFn func(answer *models.Answer) (*models.Answer, error)
	DeleteAnswerFn func(id string) error
	GetAnswerByIDFn func(id string) (*models.Answer, error)
	GetAllQuestionAnswersFn func(questionId string) ([]*models.Answer, error)
)

func (o *fakeQuestionOptionService) CreateAnswer(answer *models.Answer) (*models.Answer, error){
	return CreateAnswerFn(answer)
}

func (o *fakeQuestionOptionService) UpdateAnswer(answer *models.Answer) (*models.Answer, error) {
	return UpdateAnswerFn(answer)
}

func (o *fakeQuestionOptionService) DeleteAnswer(id string) error {
	return DeleteAnswerFn(id)
}

func (o *fakeQuestionOptionService) GetAnswerByID(id string) (*models.Answer, error) {
	return GetAnswerByIDFn(id)
}

func (o *fakeQuestionOptionService) GetAllQuestionAnswers(questionId string) ([]*models.Answer, error) {
	return GetAllQuestionAnswersFn(questionId)
}





