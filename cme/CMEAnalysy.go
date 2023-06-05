package cme

type CMEAnalysy struct {
	StartDate        string `json:"startDate"`
	EndDate          string `json:"endDate"`
	MostAccurateOnly string `json:"mostAccurateOnly"`
	Speed            string `json:"speed"`
	HalfAngle        string `json:"halfAngle"`
	Catalog          string `json:"catalog"`
}
