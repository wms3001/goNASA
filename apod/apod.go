package apod

import "github.com/wms3001/goNASA/base"

type ApdReq struct {
	Date       string `json:"date"`
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
	Count      string `json:"count"`
	Thumbs     string `json:"thumbs"`
}

type ApdRes struct {
	base.NASARsp
	Copyright   string `json:"copyright"`
	Date        string `json:"date"`
	Explanation string `json:"explanation"`
	Hdurl       string `json:"hdurl"`
	MediaType   string `json:"media_type"`
	Title       string `json:"title"`
	Url         string `json:"url"`
}

type Apod struct {
	AdReq ApdReq `json:"adReq"`
	AdRes ApdRes `json:"adRes"`
}
