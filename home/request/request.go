package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	userId := 12

	deleteURL := fmt.Sprintf("http://localhost:8080/delete_resource?user_id=%d", userId)
	deleteReq, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	deleteResp, err := client.Do(deleteReq)
	if err != nil {
		log.Fatal(err)
	}
	defer deleteResp.Body.Close()

	if deleteResp.StatusCode != http.StatusOK {
		log.Fatalf("DELETE request failed with status code: %d", deleteResp.StatusCode)
	}
	fmt.Println("DELETE request successful")

	getURL := "http://localhost:8080/get_resource"
	getReq, err := http.NewRequest("GET", getURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	getResp, err := client.Do(getReq)
	if err != nil {
		log.Fatal(err)
	}
	defer getResp.Body.Close()

	if getResp.StatusCode != http.StatusOK {
		log.Fatalf("GET request failed with status code: %d", getResp.StatusCode)
	}

	getBody, err := io.ReadAll(getResp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET request successful")
	fmt.Println("Response Body:", string(getBody))
}
