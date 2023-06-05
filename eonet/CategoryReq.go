package eonet

type CategoryReq struct {
	CategoryId string `json:"categoryId"`
	Status     string `json:"status"`
	Limit      string `json:"limit"`
	Days       string `json:"days"`
	Source     string `json:"source"`
}
