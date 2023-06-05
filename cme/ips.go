package cme

type IPS struct {
	Catalog     string `json:"catalog"`
	ActivityID  string `json:"activityID"`
	Location    string `json:"location"`
	EventTime   string `json:"eventTime"`
	Link        string `json:"link"`
	Instruments []struct {
		DisplayName string `json:"displayName"`
	} `json:"instruments"`
}

type IPSReq struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Location  string `json:"location"`
	Catalog   string `json:"catalog"`
}
