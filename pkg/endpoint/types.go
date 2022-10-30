package endpoint

type CreateUserRequest struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Err      string `json:"err"`
}

type DeleteUserRequest struct {
	Email string `json:"email"`
}

type DeleteUserResponse struct {
	Ok  bool   `json:"ok"`
	Err string `json:"err"`
}

type GetUserRequest struct {
	Email string `json:"email"`
}

type GetUserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Err      string `json:"err"`
}

type UpdateUserRequest struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type UpdateUserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Err      string `json:"err"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
	Err   string `json:"err"`
}
