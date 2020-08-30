package interfaces

import (
	"multi-choice/app/domain/repository/answer"
	"multi-choice/app/domain/repository/question"
	"multi-choice/app/domain/repository/question_option"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AnsService            answer.AnsService
	QuestionService       question.QuesService
	//QuestionService       mock.QuesService
	QuestionOptionService question_option.OptService
}

