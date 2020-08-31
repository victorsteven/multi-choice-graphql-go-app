package interfaces_test

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"
	"multi-choice/app/domain/repository/answer"
	"multi-choice/app/generated"
	"multi-choice/app/interfaces"
	"multi-choice/app/models"
	"testing"
)

type fakeAnswerService struct {}

var fakeAnswer answer.AnsService = &fakeAnswerService{} //this is where the real implementation is swap with our fake implementation

func TestCreateAnswer_Success(t *testing.T) {

	var srv = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		AnsService: fakeAnswer,  //this is swap with the real interface
		QuestionOptionService: fakeQuestionOption, //this is swap with the real interface
	}})))

	//We dont call the domain method, we swap it with this
	CreateAnswerFn = func(answer *models.Answer) (*models.Answer, error) {
		return &models.Answer{
			ID: "1",
			QuestionID: "1",
			OptionID: "1",
			IsCorrect: false,
		}, nil
	}

	//We dont call the domain method, we swap it with this
	GetQuestionOptionByIDFn = func(id string) (*models.QuestionOption, error) {
		return &models.QuestionOption{
			ID: "1",
			Title: "Option 1",
			IsCorrect: true,
		}, nil
	}

	var resp struct {
		CreateAnswer struct {
			Message     string
			Status      int
			Data       models.Answer
		}
	}

	srv.MustPost(`mutation { CreateAnswer(questionId: "1", optionId: "1") { message, status, data { id questionId optionId isCorrect } }}`, &resp)

	assert.Equal(t, 201, resp.CreateAnswer.Status)
	assert.Equal(t, "Successfully created answer", resp.CreateAnswer.Message)
	assert.Equal(t, false, resp.CreateAnswer.Data.IsCorrect)
	assert.Equal(t, "1", resp.CreateAnswer.Data.ID)
}


func TestUpdateAnswer_Success(t *testing.T) {

	var srv = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		AnsService: fakeAnswer, //this is swap with the real interface
	}})))

	//We dont call the domain method, we swap it with this
	GetAnswerByIDFn = func(id string) (*models.Answer, error) {
		return &models.Answer{
			ID: "1",
			QuestionID: "1",
			OptionID: "1",
		}, nil
	}

	//We dont call the domain method, we swap it with this
	UpdateAnswerFn = func(ans *models.Answer) (*models.Answer, error) {
		return &models.Answer{
			ID: "1",
			QuestionID: "1",
			OptionID: "2",
			IsCorrect: false,
		}, nil
	}

	var resp struct {
		UpdateAnswer struct {
			Message     string
			Status      int
			Data       	models.Answer
		}
	}

	srv.MustPost(`mutation { UpdateAnswer(id: "1", questionId: "1", optionId: "2") { message, status, data { id questionId optionId isCorrect } }}`, &resp)

	assert.Equal(t, 200, resp.UpdateAnswer.Status)
	assert.Equal(t, "Successfully updated answer", resp.UpdateAnswer.Message)
	assert.Equal(t, "1", resp.UpdateAnswer.Data.ID)
	assert.Equal(t, false, resp.UpdateAnswer.Data.IsCorrect)
}


func TestDeleteAnswer_Success(t *testing.T) {

	var srv = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		AnsService: fakeAnswer, //this is swap with the real interface
	}})))

	//We dont call the domain method, we swap it with this
	DeleteAnswerFn = func(id string)  error {
		return nil
	}

	var resp struct {
		DeleteAnswer struct {
			Message     string
			Status      int
		}
	}

	srv.MustPost(`mutation { DeleteAnswer(id: "1") { message, status }}`, &resp)

	assert.Equal(t, 200, resp.DeleteAnswer.Status)
	assert.Equal(t, "Successfully deleted answer", resp.DeleteAnswer.Message)
}


func TestGetOneAnswer_Success(t *testing.T) {

	var srv = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		AnsService: fakeAnswer, //this is swap with the real interface
	}})))

	//We dont call the domain method, we swap it with this
	GetAnswerByIDFn = func(id string) (*models.Answer, error) {
		return &models.Answer{
			ID: "1",
			QuestionID: "1",
			OptionID: "1",
			IsCorrect: true,
		}, nil
	}

	var resp struct {
		GetOneAnswer struct {
			Message     string
			Status      int
			Data       models.Answer
		}
	}

	srv.MustPost(`query { GetOneAnswer(id: "1") { 
			message, status, data { id questionId optionId isCorrect } 

	}}`, &resp)

	assert.Equal(t, 200, resp.GetOneAnswer.Status)
	assert.Equal(t, "Successfully retrieved answer", resp.GetOneAnswer.Message)
	assert.Equal(t, true, resp.GetOneAnswer.Data.IsCorrect)
	assert.Equal(t, "1", resp.GetOneAnswer.Data.ID)
	assert.Equal(t, "1", resp.GetOneAnswer.Data.QuestionID)
	assert.Equal(t, "1", resp.GetOneAnswer.Data.OptionID)
}

func TestGetAllQuestionAnswers_Success(t *testing.T) {

	var srv = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		AnsService: fakeAnswer,  //this is swap with the real interface
	}})))

	//We dont call the domain method, we swap it with this
	GetAllQuestionAnswersFn = func(questionId string) ([]*models.Answer, error) {
		return []*models.Answer{
			{
				ID: "1",
				QuestionID: "1",
				OptionID: "1",
				IsCorrect: true,
			},
			{
				ID: "2",
				QuestionID: "1",
				OptionID: "2",
				IsCorrect: false,
			},
		}, nil
	}

	var resp struct {
		GetAllQuestionAnswers struct {
			Message     string
			Status      int
			DataList       []*models.Answer
		}
	}

	srv.MustPost(`query { GetAllQuestionAnswers(questionId: "1") { 
			message, status, dataList { id questionId optionId isCorrect } 
	}}`, &resp)

	assert.Equal(t, 200, resp.GetAllQuestionAnswers.Status)
	assert.Equal(t, "Successfully retrieved all answers", resp.GetAllQuestionAnswers.Message)
	assert.Equal(t, 2, len(resp.GetAllQuestionAnswers.DataList))
}
