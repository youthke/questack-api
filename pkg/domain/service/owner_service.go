package service

import (
	"github.com/HazeyamaLab/questack-api/pkg/domain/model"
	"github.com/HazeyamaLab/questack-api/pkg/repository"
	"github.com/HazeyamaLab/questack-api/pkg/util"
	"log"
)

type OwnerService interface {
	Create(name string, mail string, password string) (uint,error)
}

type ownerService struct {
	 ownerRepository repository.OwnerRepository

}


func NewOwnerService(repository repository.OwnerRepository) OwnerService{
	return &ownerService{repository}
}


func (o *ownerService) Create(name string, mail string, password string) (uint, error){

	hash, err := util.GenerateHash(password)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	owner := model.Owner{
		Name: name,
		Mail: mail,
		Password: hash,
	}
	return o.ownerRepository.Create(owner)
}
