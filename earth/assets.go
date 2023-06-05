package earth

type Assets struct {
	Date     string `json:"date"`
	Id       string `json:"id"`
	Resource struct {
		Dataset string `json:"dataset"`
		Planet  string `json:"planet"`
	} `json:"resource"`
	ServiceVersion string `json:"service_version"`
	Url            string `json:"url"`
}
