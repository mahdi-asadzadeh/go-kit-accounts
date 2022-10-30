package types

type CreateUserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type DeleteUserResponse struct {
	Ok bool `json:"ok"`
}

type GetUserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type UpdateUserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type ErrorResponse struct {
	StatusCode int         `json:"statuscode"`
	Method     string      `json:"method"`
	Error      interface{} `json:"error"`
}
