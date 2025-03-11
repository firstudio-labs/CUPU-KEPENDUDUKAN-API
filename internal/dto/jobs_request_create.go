package dto

type JobReqCreate struct {
	Code string `validate:"required,min=1" json:"code"`
	Name string `validate:"required,min=1" json:"name"`
}
