package cme

type RBE struct {
	RbeID       string `json:"rbeID"`
	EventTime   string `json:"eventTime"`
	Instruments []struct {
		DisplayName string `json:"displayName"`
	} `json:"instruments"`
	LinkedEvents []struct {
		ActivityID string `json:"activityID"`
	} `json:"linkedEvents"`
	Link string `json:"link"`
}
