package requests

type CreateUserReq struct {
	FirstName  string  `validate:"required,min=1,max=100" json:"first_name"`
	LastName   string  `validate:"required,min=1,max=100" json:"last_name"`
	Email      string  `validate:"required,email" json:"email"`
	Password   string  `validate:"required" json:"password"`
	Occupation *string `json:"occupation"`
	Address    *string `json:"address"`
}

type LoginReq struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}
