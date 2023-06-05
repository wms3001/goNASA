package goNASA

type Teah struct {
	Results [][]interface{} `json:"results"`
	Count   int             `json:"count"`
	Total   int             `json:"total"`
	Perpage int             `json:"perpage"`
	Page    int             `json:"page"`
}
