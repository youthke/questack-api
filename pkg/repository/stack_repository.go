package repository

import (
	"github.com/HazeyamaLab/questack-api/pkg/domain/model"
	"github.com/jinzhu/gorm"
)

type StackRepository interface {
	Create(stack model.Stack) error
	FindAllByOwnerID(id uint)([]model.Stack, error)
	FindOne(id uint)(model.Stack, error)
	FindOneByID(url string)(model.Stack, error)
	Update(stack model.Stack)error
}

type stackRepository struct {
	db *gorm.DB
}

func NewStackRepository(db *gorm.DB)StackRepository{
	return &stackRepository{db: db}
}


func(s *stackRepository)Create(stack model.Stack) error{
	return s.db.Create(&stack).Error
}

func(s *stackRepository)FindAllByOwnerID(id uint)([]model.Stack, error){
	var stacks []model.Stack
	err := s.db.Find(&stacks,"owner_id=?", id).Error
	return stacks, err
}

func(s *stackRepository)FindOne(id uint)(model.Stack, error){
	var stack model.Stack
	err := s.db.First(&stack, "id = ?", id).Error
	return stack, err
}

func(s *stackRepository)Update(stack model.Stack) error{
	return s.db.Save(&stack).Error
}

func(s *stackRepository)FindOneByID(id string)(model.Stack ,error){
	var stack model.Stack
	err := s.db.Debug().Preload("Questions").First(&stack,"id=?",id).Error
	return stack, err
}

//func(s *stackRepository)ExistCheck(url string) (uint, error){
//	var stack model.Stack
//	err := s.db.First(&stack,"url=?",url).Error
//	if err != nil {
//		return 0, err
//	}
//	return stack.ID, nil
//}
