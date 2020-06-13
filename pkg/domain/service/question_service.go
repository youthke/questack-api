package service

import (
	"github.com/HazeyamaLab/questack-api/pkg/domain/model"
	"github.com/HazeyamaLab/questack-api/pkg/repository"
)

type QuestionService interface {
	Create(stackURL string, title string, author string, content string) error
}

type questionService struct{
	repository.QuestionRepository
	repository.StackRepository
}

func NewQuestionService(r repository.QuestionRepository, s repository.StackRepository) QuestionService{
	return &questionService{r,s}
}


func (q *questionService)Create(stackURL string, title string, author string, content string) error{
	stackID, err := q.StackRepository.ExistCheck(stackURL)
	if err !=  nil{
		return err
	}
	question := model.Question{
		Title:title,
		Author:author,
		Content:content,
		StackRefer:stackID,
	}
	return q.QuestionRepository.Create(question)
}