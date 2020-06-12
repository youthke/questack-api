package repository

import (
	"github.com/HazeyamaLab/questack-api/pkg/domain/model"
	"github.com/jinzhu/gorm"
)

type OwnerRepository interface {
	Create(owner model.Owner)(uint, error)
}

type ownerRepository struct {
	db *gorm.DB
}

func NewOwnerRepository(db *gorm.DB) OwnerRepository{
	return &ownerRepository{db: db}
}

func (u *ownerRepository)Create(owner model.Owner) (uint,error){
	err :=  u.db.Create(&owner).Error
	return owner.ID, err
}
