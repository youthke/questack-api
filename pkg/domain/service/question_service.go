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
}

func NewQuestionService(r repository.QuestionRepository) QuestionService{
	return &questionService{r}
}


func (q *questionService)Create(stackID string, title string, author string, content string) error{

	question := model.Question{
		Title:title,
		Author:author,
		Content:content,
		StackRefer:stackID,
	}
	return q.QuestionRepository.Create(question)
}