package presenter

type QuestionCreateForm struct{
	StackURL string `json:"stack_url"`
	Title string `json:"title"`
	Author string `json:"author"`
	Content string `json:"content"`
}
