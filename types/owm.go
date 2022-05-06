package types

type OWMResponse struct {
	Coord    Coordinates `json:"coord"`
	Weather  []Weather   `json:"weather"`
	Main     Main        `json:"main"`
	Other    Sys         `json:"sys"`
	Name     string      `json:"name"`
	Timezone int         `json:"timezone"`
	Dt       int64       `json:"dt"`
}

type Coordinates struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Main struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  float32 `json:"pressure"`
	Humidity  float32 `json:"humidity"`
}

type Sys struct {
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
	Country string `json:"country"`
}
