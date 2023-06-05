package cme

type GST struct {
	GstID      string `json:"gstID"`
	StartTime  string `json:"startTime"`
	AllKpIndex []struct {
		ObservedTime string  `json:"observedTime"`
		KpIndex      float64 `json:"kpIndex"`
		Source       string  `json:"source"`
	} `json:"allKpIndex"`
	LinkedEvents []struct {
		ActivityID string `json:"activityID"`
	} `json:"linkedEvents"`
	Link string `json:"link"`
}
