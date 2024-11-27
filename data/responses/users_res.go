package responses

type UserResponse struct {
	ID         int     `json:"id"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	Occupation *string `json:"occupation"`
	Address    *string `json:"address"`
}
