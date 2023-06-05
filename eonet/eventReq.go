package eonet

type EventReq struct {
	Status string `json:"status"`
	Limit  string `json:"limit"`
	Days   string `json:"days"`
	Source string `json:"source"`
}
