package donki

type CMEActive struct {
	ActivityID      string      `json:"activityID"`
	Catalog         string      `json:"catalog"`
	StartTime       string      `json:"startTime"`
	SourceLocation  string      `json:"sourceLocation"`
	ActiveRegionNum interface{} `json:"activeRegionNum"`
	Link            string      `json:"link"`
	Note            string      `json:"note"`
	Instruments     []struct {
		DisplayName string `json:"displayName"`
	} `json:"instruments"`
	CmeAnalyses []struct {
		Time215        string  `json:"time21_5"`
		Latitude       float64 `json:"latitude"`
		Longitude      float64 `json:"longitude"`
		HalfAngle      float64 `json:"halfAngle"`
		Speed          float64 `json:"speed"`
		Type           string  `json:"type"`
		IsMostAccurate bool    `json:"isMostAccurate"`
		Note           string  `json:"note"`
		LevelOfData    int     `json:"levelOfData"`
		Link           string  `json:"link"`
		EnlilList      []struct {
			ModelCompletionTime       string      `json:"modelCompletionTime"`
			Au                        float64     `json:"au"`
			EstimatedShockArrivalTime interface{} `json:"estimatedShockArrivalTime"`
			EstimatedDuration         interface{} `json:"estimatedDuration"`
			RminRe                    interface{} `json:"rmin_re"`
			Kp18                      interface{} `json:"kp_18"`
			Kp90                      interface{} `json:"kp_90"`
			Kp135                     interface{} `json:"kp_135"`
			Kp180                     interface{} `json:"kp_180"`
			IsEarthGB                 bool        `json:"isEarthGB"`
			Link                      string      `json:"link"`
			ImpactList                []struct {
				IsGlancingBlow bool   `json:"isGlancingBlow"`
				Location       string `json:"location"`
				ArrivalTime    string `json:"arrivalTime"`
			} `json:"impactList"`
			CmeIDs []string `json:"cmeIDs"`
		} `json:"enlilList"`
	} `json:"cmeAnalyses"`
	LinkedEvents interface{} `json:"linkedEvents"`
}

type CME struct {
	CMEActives []CMEActive `json:"cMEActives"`
}
