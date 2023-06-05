package cme

type SEP struct {
	SepID       string `json:"sepID"`
	EventTime   string `json:"eventTime"`
	Instruments []struct {
		DisplayName string `json:"displayName"`
	} `json:"instruments"`
	LinkedEvents []struct {
		ActivityID string `json:"activityID"`
	} `json:"linkedEvents"`
	Link string `json:"link"`
}
