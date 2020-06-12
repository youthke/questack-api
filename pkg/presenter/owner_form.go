package presenter

type OwnerCreateForm struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
	Password string `json:"password"`
}

