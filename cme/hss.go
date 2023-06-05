package cme

type HSS struct {
	HssID       string `json:"hssID"`
	EventTime   string `json:"eventTime"`
	Instruments []struct {
		DisplayName string `json:"displayName"`
	} `json:"instruments"`
	LinkedEvents interface{} `json:"linkedEvents"`
	Link         string      `json:"link"`
}
