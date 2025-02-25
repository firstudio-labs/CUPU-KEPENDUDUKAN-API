package dto

type JobReqCreate struct {
	Code string `validate:"required,min=5" json:"code"`
	Name string `validate:"required,min=5" json:"name"`
}
