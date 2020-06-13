package presenter

type StackCreateForm struct {
	Name string
}


type StackUpdateForm struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}