package model

type Owner struct {
	ID uint `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Mail string `json:"mail"`
	Password string `json:"-"`
	Stacks []Stack `gorm:"foreignkey:OwnerRefer" json:"stacks"`
}


