package dto

type JobReqUpdate struct {
	Code string `validate:"min=1" json:"code"`
	Name string `json:"name"`
}
