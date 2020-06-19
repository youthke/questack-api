package service

import (
	"errors"
	"github.com/HazeyamaLab/questack-api/pkg/domain/model"
	"github.com/HazeyamaLab/questack-api/pkg/repository"
	"github.com/HazeyamaLab/questack-api/pkg/util"
)

type StackService interface {
	Create(name string, ownerID uint) error
	FindALLByOwnerID(id uint)([]model.Stack,error)
	Update(stackID uint, name string, ownerID uint) error
	FindOneByID(id string)(model.Stack, error)
}

type stackService struct {
	repository.StackRepository
}

func NewStackService(s repository.StackRepository) StackService{
	return &stackService{s}
}

func(s *stackService)Create(name string, ownerID uint) error{
	id := util.RandString()
	stack := model.Stack{
		ID: id,
		Name:name,
		OwnerRefer:ownerID,
	}
	return s.StackRepository.Create(stack)
}


func(s *stackService)FindALLByOwnerID(id uint) ([]model.Stack,error){
	return s.StackRepository.FindAllByOwnerID(id)
}

func (s *stackService)FindOneByID(id string)(model.Stack, error){
	return s.StackRepository.FindOneByID(id)
}

func(s *stackService)Update(stackID uint, name string, ownerID uint) error{

	stack , err := s.StackRepository.FindOne(stackID)
	if err != nil {
		return err
	}
	if stack.OwnerRefer != ownerID{
		return errors.New("unauthorized")
	}

	stack.SetName(name)

	return s.StackRepository.Update(stack)
}



