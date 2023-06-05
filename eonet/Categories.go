package eonet

type CategoriesInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Categories  []struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Link        string `json:"link"`
		Description string `json:"description"`
		Layers      string `json:"layers"`
	} `json:"categories"`
}
