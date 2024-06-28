package models

type WeatherData struct {
	CurrentCondition []struct {
		FeelsLikeC  string `json:"FeelsLikeC"`
		Humidity    string `json:"humidity"`
		TempC       string `json:"temp_C"`
		WeatherDesc []struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
		WindSpeedKmph string `json:"windspeedKmph"`
	} `json:"current_condition"`
	Weather []struct {
		Date     string `json:"date"`
		MaxtempC string `json:"maxtempC"`
		MintempC string `json:"mintempC"`
		Hourly   []struct {
			Time          string `json:"time"`
			TempC         string `json:"tempC"`
			WindSpeedKmph string `json:"windspeedKmph"`
		} `json:"hourly"`
	} `json:"weather"`
}
