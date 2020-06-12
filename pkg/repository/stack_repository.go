package repository

import (
	"github.com/HazeyamaLab/questack-api/pkg/domain/model"
	"github.com/jinzhu/gorm"
)

type StackRepository interface {
	Create(stack model.Stack) error
	FindAllByOwnerID(id uint)([]model.Stack, error)
}

type stackRepository struct {
	db *gorm.DB
}

func NewStackRepository(db *gorm.DB)StackRepository{
	return &stackRepository{db: db}
}


func(s *stackRepository)Create(stack model.Stack) error{
	return s.db.Debug().Create(&stack).Error
}

func(s *stackRepository)FindAllByOwnerID(id uint)([]model.Stack, error){
	var stacks []model.Stack
	err := s.db.First(&stacks,"owner_id=?", id).Error
	return stacks, err
}
