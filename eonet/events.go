package eonet

import "time"

type Events struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Events      []struct {
		Id          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Link        string `json:"link"`
		Categories  []struct {
			Id    int    `json:"id"`
			Title string `json:"title"`
		} `json:"categories"`
		Sources []struct {
			Id  string `json:"id"`
			Url string `json:"url"`
		} `json:"sources"`
		Geometries []struct {
			Date        time.Time `json:"date"`
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometries"`
	} `json:"events"`
}
