package epic

type Epic struct {
	Identifier          string `json:"identifier"`
	Caption             string `json:"caption"`
	Image               string `json:"image"`
	ImageUrl            string `json:"imageUrl"`
	Version             string `json:"version"`
	CentroidCoordinates struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"centroid_coordinates"`
	DscovrJ2000Position struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"dscovr_j2000_position"`
	LunarJ2000Position struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"lunar_j2000_position"`
	SunJ2000Position struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"sun_j2000_position"`
	AttitudeQuaternions struct {
		Q0 float64 `json:"q0"`
		Q1 float64 `json:"q1"`
		Q2 float64 `json:"q2"`
		Q3 float64 `json:"q3"`
	} `json:"attitude_quaternions"`
	Date   string `json:"date"`
	Coords struct {
		CentroidCoordinates struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"centroid_coordinates"`
		DscovrJ2000Position struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"dscovr_j2000_position"`
		LunarJ2000Position struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"lunar_j2000_position"`
		SunJ2000Position struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"sun_j2000_position"`
		AttitudeQuaternions struct {
			Q0 float64 `json:"q0"`
			Q1 float64 `json:"q1"`
			Q2 float64 `json:"q2"`
			Q3 float64 `json:"q3"`
		} `json:"attitude_quaternions"`
	} `json:"coords"`
}
