package model

import "github.com/jinzhu/gorm"

type Owner struct {
	gorm.Model
	ID uint `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Mail string `json:"mail"`
	Password string `json:"-"`
	Stacks []Stack `gorm:"foreignkey:OwnerRefer" json:"stacks"`
}


