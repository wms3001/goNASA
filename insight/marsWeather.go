package insight

type WarsWeather struct {
	SolKeys        []interface{} `json:"sol_keys"`
	ValidityChecks struct {
		Field1 struct {
			AT struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"AT"`
			HWS struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"HWS"`
			PRE struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"PRE"`
			WD struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"WD"`
		} `json:"1219"`
		SolHoursRequired int      `json:"sol_hours_required"`
		SolsChecked      []string `json:"sols_checked"`
	} `json:"validity_checks"`
}
