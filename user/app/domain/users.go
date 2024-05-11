package domain

type Users struct {
	Email     string `json:"email" binding:"required,email"`
	FirstName string `json:"first_name" binding:"required,min=3"`
	LastName  string `json:"last_name" binding:"required,min=3"`
	Password  string `json:"password,omitempty" binding:"required,min=8"`
	Status    bool   `json:"status"`
}
