package marsPhotos

type MarsPhotos struct {
	Photos []struct {
		Id     int `json:"id"`
		Sol    int `json:"sol"`
		Camera struct {
			Id       int    `json:"id"`
			Name     string `json:"name"`
			RoverId  int    `json:"rover_id"`
			FullName string `json:"full_name"`
		} `json:"camera"`
		ImgSrc    string `json:"img_src"`
		EarthDate string `json:"earth_date"`
		Rover     struct {
			Id          int    `json:"id"`
			Name        string `json:"name"`
			LandingDate string `json:"landing_date"`
			LaunchDate  string `json:"launch_date"`
			Status      string `json:"status"`
		} `json:"rover"`
	} `json:"photos"`
}
