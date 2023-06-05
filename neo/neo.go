package neo

import "github.com/wms3001/goNASA/base"

type NeReq struct {
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
}

type NeoLink struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Id                 string  `json:"id"`
	NeoReferenceId     string  `json:"neo_reference_id"`
	Name               string  `json:"name"`
	NasaJplUrl         string  `json:"nasa_jpl_url"`
	AbsoluteMagnitudeH float64 `json:"absolute_magnitude_h"`
	EstimatedDiameter  struct {
		Kilometers struct {
			EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
			EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
		} `json:"kilometers"`
		Meters struct {
			EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
			EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
		} `json:"meters"`
		Miles struct {
			EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
			EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
		} `json:"miles"`
		Feet struct {
			EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
			EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
		} `json:"feet"`
	} `json:"estimated_diameter"`
	IsPotentiallyHazardousAsteroid bool `json:"is_potentially_hazardous_asteroid"`
	CloseApproachData              []struct {
		CloseApproachDate      string `json:"close_approach_date"`
		CloseApproachDateFull  string `json:"close_approach_date_full"`
		EpochDateCloseApproach int64  `json:"epoch_date_close_approach"`
		RelativeVelocity       struct {
			KilometersPerSecond string `json:"kilometers_per_second"`
			KilometersPerHour   string `json:"kilometers_per_hour"`
			MilesPerHour        string `json:"miles_per_hour"`
		} `json:"relative_velocity"`
		MissDistance struct {
			Astronomical string `json:"astronomical"`
			Lunar        string `json:"lunar"`
			Kilometers   string `json:"kilometers"`
			Miles        string `json:"miles"`
		} `json:"miss_distance"`
		OrbitingBody string `json:"orbiting_body"`
	} `json:"close_approach_data"`
	IsSentryObject bool `json:"is_sentry_object"`
}

type NeoLinkList struct {
	Links struct {
		Next string `json:"next"`
		Self string `json:"self"`
	} `json:"links"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"total_elements"`
		TotalPages    int `json:"total_pages"`
		Number        int `json:"number"`
	} `json:"page"`
	Near_earth_objects []NeoLink `json:"near_earth_objects"`
}

type NeRes struct {
	base.NASARsp

	Links struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Self     string `json:"self"`
	} `json:"links"`
	ElementCount       int                  `json:"element_count"`
	Near_earth_objects map[string][]NeoLink `json:"near_earth_objects"`
}

type NEO struct {
	NeoReq NeReq `json:"neoReq"`
	NeoRes NeRes `json:"neoRes"`
}
