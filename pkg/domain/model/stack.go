package model

type Stack struct{
	ID uint `gorm:"primary_key" json:"id"`
	Owner Owner `json:"owner"`
	OwnerRefer uint
	Name string `gorm:"name"`
	Questions []Question `gorm:"foreignkey:StackRefer" json:"questions"`
}


