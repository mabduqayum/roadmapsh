package model

type WeatherData struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Conditions  string  `json:"conditions"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
