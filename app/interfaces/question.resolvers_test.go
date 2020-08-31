package interfaces_test

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"
	"multi-choice/app/domain/repository/question"
	"multi-choice/app/domain/repository/question_option"
	"multi-choice/app/generated"
	"multi-choice/app/interfaces"
	"multi-choice/app/models"
	"testing"
)

type fakeQuestionService struct{}

type fakeQuestionOptionService struct{}

var fakeQuestion question.QuesService = &fakeQuestionService{} //this is where the real implementation is swap with our fake implementation

var fakeQuestionOption question_option.OptService = &fakeQuestionOptionService{} //this is where the real implementation is swap with our fake implementation

func TestCreateQuestion_Success(t *testing.T) {

	var srv = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		QuestionService:       fakeQuestion,       //this is swap with the real interface
		QuestionOptionService: fakeQuestionOption, //this is swap with the real interface
	}})))

	//We dont call the domain method, we swap it with this
	CreateQuestionFn = func(question *models.Question) (*models.Question, error) {
		return &models.Question{
			ID:    "1",
			Title: "Question title",
		}, nil
	}

	//also the mock on the question option:
	CreateQuestionOptionFn = func(question *models.QuestionOption) (*models.QuestionOption, error) {
		return &models.QuestionOption{
			ID:        "1",
			Title:     "Option 1",
			Position:  1,
			IsCorrect: false,
		}, nil
	}

	var resp struct {
		CreateQuestion struct {
			Message string
			Status  int
			Data    models.Question
		}
	}

	srv.MustPost(`mutation { CreateQuestion(question:{title:"Question title", options: [{title: "Option 1", position: 1, isCorrect: false}]}) { message, status, data { id title } }}`, &resp)

	assert.Equal(t, 201, resp.CreateQuestion.Status)
	assert.Equal(t, "Successfully created question", resp.CreateQuestion.Message)
	assert.Equal(t, "Question title", resp.CreateQuestion.Data.Title)
	assert.Equal(t, "1", resp.CreateQuestion.Data.ID)
}

func TestUpdateQuestion_Success(t *testing.T) {

	var srv = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		QuestionService:       fakeQuestion,       //this is swap with the real interface
		QuestionOptionService: fakeQuestionOption, //this is swap with the real interface
	}})))

	//We dont call the domain method, we swap it with this
	GetQuestionByIDFn = func(id string) (*models.Question, error) {
		return &models.Question{
			ID:    "1",
			Title: "Question title",
		}, nil
	}

	//We dont call the domain method, we swap it with this
	DeleteQuestionOptionByQuestionIDFn = func(questionId string) error {
		return nil
	}

	//We dont call the domain method, we swap it with this
	UpdateQuestionFn = func(question *models.Question) (*models.Question, error) {
		return &models.Question{
			ID:    "1",
			Title: "Question title updated",
		}, nil
	}

	//also the mock on the question option:
	CreateQuestionOptionFn = func(question *models.QuestionOption) (*models.QuestionOption, error) {
		return &models.QuestionOption{
			ID:        "1",
			Title:     "Option 1",
			Position:  1,
			IsCorrect: true,
		}, nil
	}

	var resp struct {
		UpdateQuestion struct {
			Message string
			Status  int
			Data    models.Question
		}
	}

	srv.MustPost(`mutation { UpdateQuestion(id: "1", question:{title:"Question title updated", options: [{title: "Option 1", position: 1, isCorrect: true}]}) { message, status, data { id title } }}`, &resp)

	assert.Equal(t, 200, resp.UpdateQuestion.Status)
	assert.Equal(t, "Successfully updated question", resp.UpdateQuestion.Message)
	assert.Equal(t, "Question title updated", resp.UpdateQuestion.Data.Title)
	assert.Equal(t, "1", resp.UpdateQuestion.Data.ID)
}

func TestDeleteQuestion_Success(t *testing.T) {

	var srv = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		QuestionService:       fakeQuestion,       //this is swap with the real interface
		QuestionOptionService: fakeQuestionOption, //this is swap with the real interface
	}})))

	//We dont call the domain method, we swap it with this
	DeleteQuestionFn = func(id string) error {
		return nil
	}

	//We dont call the domain method, we swap it with this
	DeleteQuestionOptionByQuestionIDFn = func(questionId string) error {
		return nil
	}

	var resp struct {
		DeleteQuestion struct {
			Message string
			Status  int
		}
	}

	srv.MustPost(`mutation { DeleteQuestion(id: "1") { message, status }}`, &resp)

	assert.Equal(t, 200, resp.DeleteQuestion.Status)
	assert.Equal(t, "Successfully deleted question", resp.DeleteQuestion.Message)
}

func TestGetOneQuestion_Success(t *testing.T) {

	var srv = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		QuestionService: fakeQuestion, //this is swap with the real interface
	}})))

	//just populate only one option
	options := []*models.QuestionOption{
		{
			ID:        "1",
			Title:     "Option 1",
			Position:  1,
			IsCorrect: false,
		},
	}

	//We dont call the domain method, we swap it with this
	GetQuestionByIDFn = func(id string) (*models.Question, error) {
		return &models.Question{
			ID:             "1",
			Title:          "Question title",
			QuestionOption: options,
		}, nil
	}

	var resp struct {
		GetOneQuestion struct {
			Message string
			Status  int
			Data    models.Question
		}
	}

	srv.MustPost(`query { GetOneQuestion(id: "1") { 
			message, status, data { 
								id title, questionOption {
        									title
        									position
        									isCorrect
      									}
							} 

	}}`, &resp)

	assert.Equal(t, 200, resp.GetOneQuestion.Status)
	assert.Equal(t, "Successfully retrieved question", resp.GetOneQuestion.Message)
	assert.Equal(t, "Question title", resp.GetOneQuestion.Data.Title)
	assert.Equal(t, "1", resp.GetOneQuestion.Data.ID)
	assert.Equal(t, 1, len(resp.GetOneQuestion.Data.QuestionOption))
}

func TestGetAllQuestions_Success(t *testing.T) {

	var srv = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		QuestionService: fakeQuestion, //this is swap with the real interface
	}})))

	//just populate only one option
	options := []*models.QuestionOption{
		{
			ID:        "1",
			Title:     "Option 1",
			Position:  1,
			IsCorrect: false,
		},
		{
			ID:        "2",
			Title:     "Option 2",
			Position:  2,
			IsCorrect: true,
		},
	}

	//We dont call the domain method, we swap it with this
	GetAllQuestionsFn = func() ([]*models.Question, error) {
		return []*models.Question{
			{
				ID:             "1",
				Title:          "Question title 1",
				QuestionOption: options,
			},
			{
				ID:             "2",
				Title:          "Question title 2",
				QuestionOption: options,
			},
		}, nil
	}

	var resp struct {
		GetAllQuestions struct {
			Message  string
			Status   int
			DataList []*models.Question
		}
	}

	srv.MustPost(`query { GetAllQuestions() { 
			message, status, dataList { 
								id title, questionOption {
        									title
        									position
        									isCorrect
      									}
							} 

	}}`, &resp)

	assert.Equal(t, 200, resp.GetAllQuestions.Status)
	assert.Equal(t, "Successfully retrieved all questions", resp.GetAllQuestions.Message)
	assert.Equal(t, 2, len(resp.GetAllQuestions.DataList))
	for _, v := range resp.GetAllQuestions.DataList {
		assert.Equal(t, 2, len(v.QuestionOption))
	}
}
