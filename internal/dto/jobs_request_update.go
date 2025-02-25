package dto

type JobReqUpdate struct {
	Code string `validate:"min=5" json:"code"`
	Name string `json:"name"`
}
