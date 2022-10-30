package types

type CreateUserInput struct {
	Email    string `form:"email" json:"email" binding"required"`
	FullName string `form:"full_name" json:"full_name" binding"required"`
	Password string `form:"password" json:"password" binding"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding"required"`
	Password string `json:"password" binding"required"`
}

type UpdateUserInput struct {
	Email    string `json:"email" binding"required"`
	FullName string `json:"full_name" binding"required"`
}
