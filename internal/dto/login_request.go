package dto

type LoginRequest struct {
	Username string `validate:"required,min=5" json:"username"`
	Password string `validate:"required,min=5" json:"password"`
}
