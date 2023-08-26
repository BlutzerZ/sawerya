package dto

type RegisterUserRequest struct {
	Username string `json:"username" binding:"required,min=5"`
	Email    string `json:"email" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=8"`
}

type UpdateUsernameRequest struct {
	ID       int    `json:"id" binding:"required"`
	Username string `json:"username" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=8"`
}

type UpdatePasswordRequest struct {
	ID          int    `json:"id" binding:"required"`
	OldPassword string `json:"oldpassword" binding:"required,min=8"`
	Password    string `json:"password" binding:"required,min=8"`
}
