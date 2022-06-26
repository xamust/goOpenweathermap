package openweathermap

type ResultMsg struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []List `json:"list"`
	City    City
}

type List struct {
	Dt         float64   `json:"dt"`
	Main       Main      `json:"main"`
	Weather    []Weather `json:"weather"`
	Clouds     Clouds    `json:"clouds"`
	Wind       Wind      `json:"wind"`
	Visibility int       `json:"visibility"`
	Pop        float64   `json:"pop"`
	Sys        Sys       `json:"sys"`
	DateTxt    string    `json:"dt_txt"`
}

type City struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Coord      Coord   `json:"coord"`
	Country    string  `json:"country"`
	Population float64 `json:"population"`
	timezone   float64 `json:"timezone"`
	sunrise    float64 `json:"sunrise"`
	sunset     float64 `json:"sunset"`
}

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Main struct {
	Temp        float64 `json:"temp"`
	FeelsLike   float64 `json:"feels_like"`
	TempMin     float64 `json:"temp_min"`
	TempMax     float64 `json:"temp_max"`
	Pressure    int64   `json:"pressure"`
	SeaLevel    int64   `json:"sea_level"`
	GroundLevel int64   `json:"grnd_level"`
	Humidity    int64   `json:"humidity"`
	TempKf      float64 `json:"temp_kf"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Clouds struct {
	All int `json:"all"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Sys struct {
	Pod string `json:"pod"`
}
