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
	GetAllQuestionsFn func() ([]*models.Question, error)
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

func (q *fakeQuestionService) GetAllQuestions() ([]*models.Question, error) {
	return GetAllQuestionsFn()
}

////////////////////////////////////////
//QuestionOption Domain Mocking
var (
	CreateQuestionOptionFn func(option *models.QuestionOption) (*models.QuestionOption, error)
	UpdateQuestionOptionFn func(option *models.QuestionOption) (*models.QuestionOption, error)
	DeleteQuestionOptionFn func(string) error
	DeleteQuestionOptionByQuestionIDFn func(questionId string) error
	GetQuestionOptionByIDFn func(string) (*models.QuestionOption, error)
	GetQuestionOptionByQuestionID func(questionId string) ([]*models.QuestionOption, error)
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

func (o *fakeQuestionOptionService) DeleteQuestionOptionByQuestionID(questionId string) error {
	return DeleteQuestionOptionByQuestionIDFn(questionId)
}

func (o *fakeQuestionOptionService) GetQuestionOptionByQuestionID(questionId string) ([]*models.QuestionOption, error) {
	return GetQuestionOptionByQuestionID(questionId)
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

func (a *fakeAnswerService) CreateAnswer(answer *models.Answer) (*models.Answer, error){
	return CreateAnswerFn(answer)
}

func (a *fakeAnswerService) UpdateAnswer(answer *models.Answer) (*models.Answer, error) {
	return UpdateAnswerFn(answer)
}

func (a *fakeAnswerService) DeleteAnswer(id string) error {
	return DeleteAnswerFn(id)
}

func (a *fakeAnswerService) GetAnswerByID(id string) (*models.Answer, error) {
	return GetAnswerByIDFn(id)
}

func (a *fakeAnswerService) GetAllQuestionAnswers(questionId string) ([]*models.Answer, error) {
	return GetAllQuestionAnswersFn(questionId)
}





