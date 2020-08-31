package persistence_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"multi-choice/app/infrastructure/persistence"
	"multi-choice/app/models"
	"testing"
)

func TestCreateQuestionOption_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var quesOpt = models.QuestionOption{
		QuestionID: "1",
		Title:      "Option 1",
		Position:   1,
		IsCorrect:  false,
	}

	repo := persistence.NewQuestionOption(conn)

	q, saveErr := repo.CreateQuestionOption(&quesOpt)

	assert.Nil(t, saveErr)
	assert.EqualValues(t, q.Title, "Option 1")
	assert.EqualValues(t, q.Position, 1)
	assert.EqualValues(t, q.IsCorrect, false)
}

//Cannot create a question question twice for the same question
func TestCreateQuestionOption_Failure(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	//seed question options
	_, err = seedQuestionOptions(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	var quesOpt = models.QuestionOption{
		QuestionID: "1",
		Title:      "Option 1", //this option has already been seeded
		Position:   3,
		IsCorrect:  true,
	}

	repo := persistence.NewQuestionOption(conn)
	f, saveErr := repo.CreateQuestionOption(&quesOpt)

	dbMsg := errors.New("two question options can't have the same title, position and/or the same correct answer")

	assert.Nil(t, f)
	assert.EqualValues(t, dbMsg, saveErr)
}

func TestGetQuestionOptionByID_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	quesOpt, err := seedQuestionOption(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	repo := persistence.NewQuestionOption(conn)

	q, saveErr := repo.GetQuestionOptionByID(quesOpt.ID)

	assert.Nil(t, saveErr)
	assert.EqualValues(t, q.Title, quesOpt.Title)
}

func TestGetAllQuestionOption_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	_, err = seedQuestions(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := persistence.NewQuestion(conn)
	quesOpts, getErr := repo.GetAllQuestions()

	assert.Nil(t, getErr)
	assert.EqualValues(t, len(quesOpts), 2)
}

func TestUpdateQuestionOption_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	ques, err := seedQuestion(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	//updating
	ques.Title = "question title update"

	repo := persistence.NewQuestion(conn)
	q, updateErr := repo.UpdateQuestion(ques)

	assert.Nil(t, updateErr)
	//assert.EqualValues(t, q.ID, "1")
	assert.EqualValues(t, q.Title, "question title update")
}

//Duplicate title error
func TestUpdateQuestionOption_Failure(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	questions, err := seedQuestions(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var secondQuestion models.Question

	//get the second question title
	for _, v := range questions {
		if v.ID == "1" {
			continue
		}
		secondQuestion = v
	}
	secondQuestion.Title = "First Question" //this title belongs to the first question already, so the question question cannot use it

	repo := persistence.NewQuestion(conn)
	q, updateErr := repo.UpdateQuestion(&secondQuestion)

	dbMsg := errors.New("question title already taken")

	assert.NotNil(t, updateErr)
	assert.Nil(t, q)
	assert.EqualValues(t, dbMsg, updateErr)
}

func TestDeleteQuestionOption_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	ques, err := seedQuestion(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := persistence.NewQuestion(conn)

	deleteErr := repo.DeleteQuestion(ques.ID)

	assert.Nil(t, deleteErr)
}
