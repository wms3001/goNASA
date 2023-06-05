package eonet

type Layer struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Categories  []struct {
		Id     int    `json:"id"`
		Title  string `json:"title"`
		Layers []struct {
			Name          string `json:"name"`
			ServiceUrl    string `json:"serviceUrl"`
			ServiceTypeId string `json:"serviceTypeId"`
			Parameters    []struct {
				TILEMATRIXSET string `json:"TILEMATRIXSET"`
				FORMAT        string `json:"FORMAT"`
			} `json:"parameters"`
		} `json:"layers"`
	} `json:"categories"`
}
