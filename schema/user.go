package schema

type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
