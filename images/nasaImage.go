package images

import "time"

type NASAImage struct {
	Collection struct {
		Version string `json:"version"`
		Href    string `json:"href"`
		Items   []struct {
			Href string `json:"href"`
			Data []struct {
				Description  string    `json:"description"`
				Title        string    `json:"title"`
				Photographer string    `json:"photographer"`
				Location     string    `json:"location"`
				NasaId       string    `json:"nasa_id"`
				DateCreated  time.Time `json:"date_created"`
				Keywords     []string  `json:"keywords"`
				MediaType    string    `json:"media_type"`
				Center       string    `json:"center"`
			} `json:"data"`
			Links []struct {
				Href   string `json:"href"`
				Rel    string `json:"rel"`
				Render string `json:"render"`
			} `json:"links"`
		} `json:"items"`
		Metadata struct {
			TotalHits int `json:"total_hits"`
		} `json:"metadata"`
		Links []struct {
			Rel    string `json:"rel"`
			Prompt string `json:"prompt"`
			Href   string `json:"href"`
		} `json:"links"`
	} `json:"collection"`
}
