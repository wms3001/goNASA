package cme

type MPC struct {
	MpcID       string `json:"mpcID"`
	EventTime   string `json:"eventTime"`
	Instruments []struct {
		DisplayName string `json:"displayName"`
	} `json:"instruments"`
	LinkedEvents []struct {
		ActivityID string `json:"activityID"`
	} `json:"linkedEvents"`
	Link string `json:"link"`
}
