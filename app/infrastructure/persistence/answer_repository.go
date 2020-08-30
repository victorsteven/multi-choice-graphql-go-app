package persistence

import (
	"errors"
	"github.com/jinzhu/gorm"
	"multi-choice/app/domain/repository/answer"
	"multi-choice/app/models"
)

type ansService struct {
	db *gorm.DB
}

func NewAnswer(db *gorm.DB) *ansService {
	return &ansService{
		db,
	}
}

//We implement the interface defined in the domain
var _ answer.AnsService = &ansService{}

func (s *ansService) CreateAnswer(answer *models.Answer) (*models.Answer, error) {

	//first we need to check if the ans have been entered for this question:
	oldAns, _ := s.GetAllQuestionAnswers(answer.QuestionID)
	if len(oldAns) > 0 {
		for _, v := range oldAns {
			//We cannot have two correct answers for this type of quiz
			if v.IsCorrect == true && answer.IsCorrect {
				return nil, errors.New("cannot have two correct answers for the same question")
			}
		}
	}

	err := s.db.Create(&answer).Error
	if err != nil {
		return nil, err
	}

	return answer, nil
}

func (s *ansService) UpdateAnswer(answer *models.Answer) (*models.Answer, error) {

	err := s.db.Save(&answer).Error
	if err != nil {
		return nil, err
	}

	return answer, nil

}

func (s *ansService) DeleteAnswer(id string) error {

	ans := &models.Answer{}

	err := s.db.Where("id = ?", id).Delete(ans).Error
	if err != nil {
		return err
	}

	return nil

}

func (s *ansService) GetAnswerByID(id string) (*models.Answer, error) {

	var ans = &models.Answer{}

	err := s.db.Where("id = ?", id).Take(&ans).Error
	if err != nil {
		return nil, err
	}

	return ans, nil
}

func (s *ansService) GetAllQuestionAnswers(questionId string) ([]*models.Answer, error) {

	var answers []*models.Answer

	err := s.db.Where("question_id = ?", questionId).Find(&answers).Error
	if err != nil {
		return nil, err
	}

	return answers, nil

}
