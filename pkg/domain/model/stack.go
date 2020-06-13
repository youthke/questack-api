package model

type Stack struct{
	ID uint `gorm:"primary_key" json:"id"`
	Owner Owner `json:"owner"`
	OwnerRefer uint `gorm:"column:owner_id"`
	Name string `gorm:"name"`
	Questions []Question `gorm:"foreignkey:StackRefer" json:"questions"`
}


func (s *Stack)SetName(name string){
	s.Name = name
}

