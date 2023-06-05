package cme

type CMEAnaly struct {
	Time215         string  `json:"time215"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	HalfAngle       float64 `json:"halfAngle"`
	Speed           float64 `json:"speed"`
	Type            string  `json:"type"`
	IsMostAccurate  bool    `json:"isMostAccurate"`
	AssociatedCMEID string  `json:"associatedCMEID"`
	Note            string  `json:"note"`
	Catalog         string  `json:"catalog"`
	Link            string  `json:"link"`
}
