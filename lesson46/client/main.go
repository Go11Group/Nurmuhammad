package main

import (
	"context"
	"fmt"
	"log"
	pb "new/genproto/transportService"
	p "new/genproto/weatherService"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	Menu()
}

func Menu() {
	fmt.Println(`
	1-Transport
	2-Weather`)
	var choice int
	fmt.Print("your choice>>> ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		Transport()
	case 2:
		Weather()
	}
}

func Transport() {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	gen := pb.NewTransportServiceClient(conn)
	var request int
	fmt.Println(`
	1 - Bus schedule
	2 - bus location
	3 - bus feedback
	4 - Exit`)
	fmt.Print("Enter number which request you choose>>> ")
	fmt.Scan(&request)
	switch request {
	case 1:
		GetBusSchedule(context.Background(), gen)
	case 2:
		TrackBusLocation(context.Background(), gen)
	case 3:
		ReportTrafficJam(context.Background(), gen)
	}

	Menu()
}
func Weather() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	gen := p.NewWeatherServiceClient(conn)
	var request int
	fmt.Println(`
	1 - Weather Now
	2 - Weather next day
	3 - Weather feedback
	4 - Exit`)
	fmt.Print("Enter number which request you choose>>> ")
	fmt.Scan(&request)
	switch request {
	case 1:
		GetCurrentWeather(context.Background(), gen)
	case 2:
		GetWeatherForecast(context.Background(), gen)
	case 3:
		ReportWeatherCondition(context.Background(), gen)
	}
	Menu()
}

func GetBusSchedule(ctx context.Context, gen pb.TransportServiceClient) {
	bus := pb.BusNumber{}
	fmt.Print("Enter book number>>> ")
	fmt.Scan(&bus.Busnum)

	bus2, err := gen.GetBusSchedule(ctx, &bus)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("schedule:", bus2)
	}
}

func TrackBusLocation(ctx context.Context, gen pb.TransportServiceClient) {
	bus := pb.BusNumber{}
	fmt.Print("Enter book number>>> ")
	fmt.Scan(&bus.Busnum)

	bus2, err := gen.TrackBusLocation(ctx, &bus)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("location:", bus2)
	}
}

func ReportTrafficJam(ctx context.Context, gen pb.TransportServiceClient) {
	bus := pb.ReportTrafficJamRequest{}
	fmt.Print("Enter book number>>> ")
	fmt.Scan(&bus.Busnum)
	fmt.Print("Enter feedback there>>> ")
	fmt.Scan(&bus.Feedback)

	bus2, err := gen.ReportTrafficJam(ctx, &bus)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Succes:", bus2)
	}
}

func GetCurrentWeather(ctx context.Context, gen p.WeatherServiceClient) {
	weather := p.Place{}
	fmt.Print("Enter location>>> ")
	fmt.Scan(&weather.Place)

	bus2, err := gen.GetCurrentWeather(ctx, &weather)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(bus2)
	}
}

func GetWeatherForecast(ctx context.Context, gen p.WeatherServiceClient) {
	weather := p.Place{}
	fmt.Print("Enter location>>> ")
	fmt.Scan(&weather.Place)

	bus2, err := gen.GetWeatherForecast(ctx, &weather)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(bus2)
	}
}

func ReportWeatherCondition(ctx context.Context, gen p.WeatherServiceClient) {
	weather := p.ReportWeatherConditionRequest{}
	fmt.Print("Enter location>>> ")
	fmt.Scan(&weather.Place)
	fmt.Print("Enter feedback there>>> ")
	fmt.Scan(&weather.Feedback)

	bus2, err := gen.ReportWeatherCondition(ctx, &weather)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("accepted :", bus2)
	}
}
