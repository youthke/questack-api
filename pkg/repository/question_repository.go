package repository

import (
	"github.com/HazeyamaLab/questack-api/pkg/domain/model"
	"github.com/jinzhu/gorm"
)

type QuestionRepository interface {
	Create(question model.Question) error
}

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB)QuestionRepository{
	return &questionRepository{db:db}
}


func(q *questionRepository)Create(question model.Question) error{
	return q.db.Create(&question).Error
}

