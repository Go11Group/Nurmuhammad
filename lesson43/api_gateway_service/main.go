package main

import "github.com/Go11Group/at_lesson/lesson43/api_gateway_service/api"

func main() {
	panic(api.Routes().ListenAndServe())
}
