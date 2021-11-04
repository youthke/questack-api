package presenter

type StackCreateForm struct {
	Name string
	Description string
}


type StackUpdateForm struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}