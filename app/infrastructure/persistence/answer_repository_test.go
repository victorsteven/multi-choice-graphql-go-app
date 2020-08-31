package persistence_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"multi-choice/app/infrastructure/persistence"
	"multi-choice/app/models"
	"testing"
)

func TestCreateAnswer_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var ans = models.Answer{
		QuestionID: "1",
		OptionID:   "1",
		IsCorrect:  false,
	}

	repo := persistence.NewAnswer(conn)

	a, saveErr := repo.CreateAnswer(&ans)

	assert.Nil(t, saveErr)
	assert.EqualValues(t, a.QuestionID, "1")
	assert.EqualValues(t, a.OptionID, "1")
	assert.EqualValues(t, a.IsCorrect, false)
}

func TestGetAnswerByID_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	ans, err := seedAnswer(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	repo := persistence.NewAnswer(conn)

	a, saveErr := repo.GetAnswerByID(ans.ID)

	assert.Nil(t, saveErr)
	assert.EqualValues(t, a.QuestionID, ans.QuestionID)
	assert.EqualValues(t, a.IsCorrect, ans.IsCorrect)
	assert.EqualValues(t, a.OptionID, ans.OptionID)
}

func TestGetAllAnswersForQuestion_Success(t *testing.T) {

	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	answers, err := seedAnswers(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	fmt.Println(answers)

	repo := persistence.NewAnswer(conn)

	//since is the same question from the seeded, let pick the first iteration:
	//anss, getErr := repo.GetAllQuestionAnswers(answers[0].QuestionID)
	anss, getErr := repo.GetAllQuestionAnswers("1")

	assert.Nil(t, getErr)
	assert.EqualValues(t, len(anss), 2)
}

func TestUpdateAnswer_Success(t *testing.T) {

	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	ans, err := seedAnswer(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	//updating
	ans.IsCorrect = false

	repo := persistence.NewAnswer(conn)
	a, updateErr := repo.UpdateAnswer(ans)

	assert.Nil(t, updateErr)
	assert.EqualValues(t, a.IsCorrect, ans.IsCorrect)
}

func TestDeleteAnswer_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	ans, err := seedAnswer(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := persistence.NewAnswer(conn)

	deleteErr := repo.DeleteAnswer(ans.ID)

	assert.Nil(t, deleteErr)
}
