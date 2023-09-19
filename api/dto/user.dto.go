package dto

type RegisterUserRequest struct {
	Username string `json:"username" binding:"required,min=5"`
	Email    string `json:"email" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=8"`
}

type UpdateUsernameRequest struct {
	Username string `json:"username" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=8"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"oldpassword" binding:"required,min=8"`
	Password    string `json:"password" binding:"required,min=8"`
}

type DeleteUserRequest struct {
	Password string `json:"password" binding:"required,min=8"`
}
