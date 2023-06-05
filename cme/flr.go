package cme

type FLR struct {
	FlrID       string `json:"flrID"`
	Instruments []struct {
		DisplayName string `json:"displayName"`
	} `json:"instruments"`
	BeginTime       string      `json:"beginTime"`
	PeakTime        string      `json:"peakTime"`
	EndTime         string      `json:"endTime"`
	ClassType       string      `json:"classType"`
	SourceLocation  string      `json:"sourceLocation"`
	ActiveRegionNum int         `json:"activeRegionNum"`
	LinkedEvents    interface{} `json:"linkedEvents"`
	Link            string      `json:"link"`
}
