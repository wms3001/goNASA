package images

type ImageAsset struct {
	Collection struct {
		Version string `json:"version"`
		Href    string `json:"href"`
		Items   []struct {
			Href string `json:"href"`
		} `json:"items"`
	} `json:"collection"`
}
