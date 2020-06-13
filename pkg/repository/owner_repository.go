package repository

import (
	"github.com/HazeyamaLab/questack-api/pkg/domain/model"
	"github.com/jinzhu/gorm"
)

type OwnerRepository interface {
	Create(owner model.Owner)(uint, error)
	FindOneByMail(mail string)(model.Owner, error)
}

type ownerRepository struct {
	db *gorm.DB
}

func NewOwnerRepository(db *gorm.DB) OwnerRepository{
	return &ownerRepository{db: db}
}

func (o *ownerRepository)Create(owner model.Owner) (uint,error){
	err :=  o.db.Create(&owner).Error
	return owner.ID, err
}

func (o *ownerRepository)FindOneByMail(mail string)(model.Owner, error){
	var owner model.Owner
	err := o.db.First(&owner, "mail=?", mail).Error
	return owner, err
}
