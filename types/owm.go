package types

type Weather struct {
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

type Current struct {
	Weather []struct {
		ID int `json:"id"`
	} `json:"weather"`
	Main struct {
		FeelsLike float64 `json:"feels_like"`
	} `json:"main"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Name string `json:"name"`
}

type Forecast struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []struct {
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
		} `json:"main"`
		Weather []struct {
			ID int `json:"id"`
		} `json:"weather"`
		DtTxt string `json:"dt_txt"`
	} `json:"list"`
	City struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Country  string `json:"country"`
		Timezone int    `json:"timezone"`
		Sunrise  int    `json:"sunrise"`
		Sunset   int    `json:"sunset"`
	} `json:"city"`
}
