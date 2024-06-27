package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	pb "new/genproto/weatherService"
	"new/models"
)

type server struct {
	pb.UnimplementedWeatherServiceServer
}

func NewWeatherService() *server {
	return &server{}
}

func (s *server) GetCurrentWeather(ctx context.Context, req *pb.Place) (*pb.GetCurrentWeatherResponse, error) {
	var weatherData models.WeatherData

	// Construct the API URL
	apiURL := fmt.Sprintf("https://wttr.in/%s?format=j1", url.QueryEscape(req.Place))

	// Make the request
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, err
	}
	weather := pb.GetCurrentWeatherResponse{
		Place:         req.Place,
		Humidity:      weatherData.CurrentCondition[0].Humidity,
		TempC:         weatherData.CurrentCondition[0].TempC,
		WindSpeedKmph: weatherData.CurrentCondition[0].WindSpeedKmph,
	}

	return &weather, nil
}

func (s *server) GetWeatherForecast(ctx context.Context, req *pb.Place) (*pb.GetWeatherForecastResponse, error) {
	var weatherData models.WeatherData

	// Construct the API URL
	apiURL := fmt.Sprintf("https://wttr.in/%s?format=j1", url.QueryEscape(req.Place))

	// Make the request
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, err
	}

	var weather pb.GetWeatherForecastResponse
	if len(weatherData.Weather) > 1 {
		weather = pb.GetWeatherForecastResponse{
			Place:         req.Place,
			Date:          weatherData.Weather[1].Date,
			Max:           weatherData.Weather[1].MaxtempC,
			Min:           weatherData.Weather[1].MintempC,
			TempC:         weatherData.Weather[1].Hourly[0].TempC,
			WindSpeedKmph: weatherData.Weather[1].Hourly[0].WindSpeedKmph,
		}
	} else {
		return nil, fmt.Errorf("No weather forecast available for tomorrow.")
	}

	return &weather, nil
}

func (s *server) ReportWeatherCondition(ctx context.Context, req *pb.ReportWeatherConditionRequest) (*pb.ReportWeatherConditionResponse, error) {
	if len(req.Feedback) > 0 {
		return &pb.ReportWeatherConditionResponse{IsAccepted: true}, nil
	} else {
		return &pb.ReportWeatherConditionResponse{IsAccepted: false}, nil
	}
}
