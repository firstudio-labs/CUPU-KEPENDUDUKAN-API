package dto

type RegisterRequest struct {
	Username        string `validate:"required,min=5" json:"username"`
	Password        string `validate:"required,min=5" json:"password"`
	ConfirmPassword string `validate:"required,min=5,eqfield=Password" json:"confirm_password"`
}
