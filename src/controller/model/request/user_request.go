package request

type UserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6,containsany=!@#$%^&*()_+{}:<>?"`
	Name      string `json:"name" binding:"required,min=2,max=100"`
	BirthDate string `json:"birth_date" binding:"required,datetime=02-01-2006"`
}
