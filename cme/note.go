package cme

type Note struct {
	MessageType      string `json:"messageType"`
	MessageID        string `json:"messageID"`
	MessageURL       string `json:"messageURL"`
	MessageIssueTime string `json:"messageIssueTime"`
	MessageBody      string `json:"messageBody"`
}
