package service

import (
	"github.com/HazeyamaLab/questack-api/pkg/domain/model"
	"github.com/HazeyamaLab/questack-api/pkg/repository"
)

type StackService interface {
	Create(name string, ownerID uint) error
	FindALLByOwnerID(id uint)([]model.Stack,error)
}

type stackService struct {
	repository.StackRepository
}

func NewStackService(s repository.StackRepository) StackService{
	return &stackService{s}
}

func(s *stackService)Create(name string, ownerID uint) error{
	stack := model.Stack{
		Name:name,
		OwnerRefer:ownerID,
	}
	return s.StackRepository.Create(stack)
}


func(s *stackService)FindALLByOwnerID(id uint) ([]model.Stack,error){
	return s.StackRepository.FindAllByOwnerID(id)
}


