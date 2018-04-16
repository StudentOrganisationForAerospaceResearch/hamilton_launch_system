package main

type Weather struct {
	AirTemperature   float64 `json:"airTemperature"`
	WindSpeed        float64 `json:"windSpeed"`
	WindDirection    float64 `json:"windDirection"`
	RelativeHumidity float64 `json:"relativeHumidity"`
}

var counter = 0.1

func getWeather() (Weather, error) {
	counter += 0.1
	return Weather{
		AirTemperature:   counter,
		WindSpeed:        counter * 2,
		WindDirection:    counter * 3,
		RelativeHumidity: counter * 4,
	}, nil
}
