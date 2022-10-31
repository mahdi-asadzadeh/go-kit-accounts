package endpoint

type CreateUserRequest struct {
	Email    string `form:"email" json:"email"`
	FullName string `form:"full_name" json:"full_name"`
	Password string `form:"password" json:"password"`
}

type CreateUserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Err      string `json:"err"`
}

type DeleteUserRequest struct {
	Email string `form:"email" json:"email"`
}

type DeleteUserResponse struct {
	Ok  bool   `json:"ok"`
	Err string `json:"err"`
}

type GetUserRequest struct {
	Email string `form:"email" json:"email"`
}

type GetUserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Err      string `json:"err"`
}

type UpdateUserRequest struct {
	Email    string `form:"email" json:"email"`
	FullName string `form:"full_name" json:"full_name"`
}

type UpdateUserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Err      string `json:"err"`
}

type LoginUserRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
	Err   string `json:"err"`
}
