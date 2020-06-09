package model

type Question struct {
	ID uint `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Content string `json:"content"`
	StackRefer uint
}


