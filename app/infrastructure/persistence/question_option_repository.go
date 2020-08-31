package persistence

import (
	"errors"
	"github.com/jinzhu/gorm"
	"multi-choice/app/domain/repository/question_option"
	"multi-choice/app/models"
)

type optService struct {
	db *gorm.DB
}

func NewQuestionOption(db *gorm.DB) *optService {
	return &optService{
		db,
	}
}

//We implement the interface defined in the domain
var _ question_option.OptService = &optService{}

func (s *optService) CreateQuestionOption(questOpt *models.QuestionOption) (*models.QuestionOption, error) {

	//check if this question option title or the position or the correctness already exist for the question
	oldOpts, _ := s.GetQuestionOptionByQuestionID(questOpt.QuestionID)
	if len(oldOpts) > 0 {
		for _, v := range oldOpts {
			if v.Title == questOpt.Title || v.Position == questOpt.Position || (v.IsCorrect == true && questOpt.IsCorrect == true) {
				return nil, errors.New("two question options can't have the same title, position and/or the same correct answer")
			}
		}
	}

	err := s.db.Create(&questOpt).Error
	if err != nil {
		return nil, err
	}

	return questOpt, nil
}

func (s *optService) UpdateQuestionOption(questOpt *models.QuestionOption) (*models.QuestionOption, error) {

	err := s.db.Save(&questOpt).Error
	if err != nil {
		return nil, err
	}

	return questOpt, nil

}

func (s *optService) DeleteQuestionOption(id string) error {

	err := s.db.Delete(id).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *optService) DeleteQuestionOptionByQuestionID(questId string) error {

	err := s.db.Delete(questId).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *optService) GetQuestionOptionByID(id string) (*models.QuestionOption, error) {

	quesOpt := &models.QuestionOption{}

	err := s.db.Where("id = ?", id).Take(&quesOpt).Error
	if err != nil {
		return nil, err
	}

	return quesOpt, nil

}

func (s *optService) GetQuestionOptionByQuestionID(id string) ([]*models.QuestionOption, error) {

	var quesOpts []*models.QuestionOption

	err := s.db.Where("question_id = ?", id).Find(&quesOpts).Error
	if err != nil {
		return nil, err
	}

	return quesOpts, nil
}
