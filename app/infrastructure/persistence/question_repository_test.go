package persistence_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"multi-choice/app/infrastructure/persistence"
	"multi-choice/app/models"
	"testing"
)

func TestCreateQuestion_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var ques = models.Question{}
	ques.Title = "Good question"

	repo := persistence.NewQuestion(conn)

	q, saveErr := repo.CreateQuestion(&ques)
	assert.Nil(t, saveErr)
	assert.EqualValues(t, q.Title, "Good question")
}

func TestCreateQuestion_Failure(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	//seed the question
	_, err = seedQuestion(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	var ques = models.Question{}
	ques.Title = "First Question"

	repo := persistence.NewQuestion(conn)
	f, saveErr := repo.CreateQuestion(&ques)

	dbMsg := errors.New("question title already taken")

	assert.Nil(t, f)
	assert.EqualValues(t, dbMsg, saveErr)
}

func TestGetQuestionByID_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	ques, err := seedQuestion(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := persistence.NewQuestion(conn)

	q, saveErr := repo.GetQuestionByID(ques.ID)

	assert.Nil(t, saveErr)
	assert.EqualValues(t, q.Title, ques.Title)
}

func TestGetAllQuestion_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	_, err = seedQuestions(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := persistence.NewQuestion(conn)
	questions, getErr := repo.GetAllQuestions()

	assert.Nil(t, getErr)
	assert.EqualValues(t, len(questions), 2)
}

func TestUpdateQuestion_Success(t *testing.T) {
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
func TestUpdateQuestion_Failure(t *testing.T) {
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

func TestDeleteQuestion_Success(t *testing.T) {
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
