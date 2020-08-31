package persistence_test

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"multi-choice/app/models"
	"os"
)

func DBConn() (*gorm.DB, error) {

	if _, err := os.Stat("./../../../.env"); !os.IsNotExist(err) {
		var err error
		err = godotenv.Load(os.ExpandEnv("./../../../.env"))
		if err != nil {
			log.Fatalf("Error getting env %v\n", err)
		} else {
			fmt.Println("we have the env")
		}
		return LocalDatabase()
	}
	return CIBuild()
}

//Circle CI DB
func CIBuild() (*gorm.DB, error) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5432", "postgres", "multi-choice-test", "password")
	conn, err := gorm.Open("postgres", DBURL)
	if err != nil {
		log.Fatal("This is the error:", err)
	}
	return conn, nil
}

//Local DB
func LocalDatabase() (*gorm.DB, error) {

	dbdriver := os.Getenv("TEST_DB_DRIVER")
	host := os.Getenv("TEST_DB_HOST")
	password := os.Getenv("TEST_DB_PASSWORD")
	user := os.Getenv("TEST_DB_USER")
	dbname := os.Getenv("TEST_DB_NAME")
	port := os.Getenv("TEST_DB_PORT")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	conn, err := gorm.Open(dbdriver, DBURL)
	if err != nil {
		return nil, err
	} else {
		log.Println("CONNECTED TO: ", dbdriver)
	}

	err = conn.DropTableIfExists(&models.Question{}, &models.Answer{}, &models.QuestionOption{}).Error
	if err != nil {
		return nil, err
	}
	err = conn.Debug().AutoMigrate(
		models.Question{},
		models.Answer{},
		models.QuestionOption{},
	).Error
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func seedQuestion(db *gorm.DB) (*models.Question, error) {
	question := &models.Question{
		ID:    "1",
		Title: "First Question",
	}
	err := db.Create(&question).Error
	if err != nil {
		return nil, err
	}
	return question, nil
}

func seedQuestions(db *gorm.DB) ([]models.Question, error) {
	questions := []models.Question{
		{
			ID:    "1",
			Title: "First Question",
		},
		{
			ID:    "2",
			Title: "Second Question",
		},
	}
	for _, v := range questions {
		err := db.Create(&v).Error
		if err != nil {
			return nil, err
		}
	}
	return questions, nil
}

func seedQuestionOption(db *gorm.DB) (*models.QuestionOption, error) {
	quesOpt := &models.QuestionOption{
		ID:         "1",
		QuestionID: "1",
		Title:      "Option 1",
		Position:   1,
		IsCorrect:  false,
	}
	err := db.Create(&quesOpt).Error
	if err != nil {
		return nil, err
	}
	return quesOpt, nil
}

func seedQuestionOptions(db *gorm.DB) ([]models.QuestionOption, error) {
	quesOpts := []models.QuestionOption{
		{
			ID:         "1",
			QuestionID: "1",
			Title:      "Option 1",
			Position:   1,
			IsCorrect:  false,
		},
		{
			ID:         "2",
			QuestionID: "2",
			Title:      "Option 2",
			Position:   2,
			IsCorrect:  true,
		},
	}
	for _, v := range quesOpts {
		err := db.Create(&v).Error
		if err != nil {
			return nil, err
		}
	}
	return quesOpts, nil
}

func seedAnswer(db *gorm.DB) (*models.Answer, error) {
	ans := &models.Answer{
		QuestionID: "1",
		OptionID:   "1",
		IsCorrect:  true,
	}
	err := db.Create(&ans).Error
	if err != nil {
		return nil, err
	}
	return ans, nil
}

func seedAnswers(db *gorm.DB) ([]models.Answer, error) {
	answers := []models.Answer{
		{
			QuestionID: "1",
			OptionID:   "1",
			IsCorrect:  false,
		},
		{
			QuestionID: "1",
			OptionID:   "2",
			IsCorrect:  true,
		},
	}
	for _, v := range answers {
		err := db.Create(&v).Error
		if err != nil {
			return nil, err
		}
	}

	return answers, nil
}
