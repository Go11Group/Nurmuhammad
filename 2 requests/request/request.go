package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	userId := "a44afced-1626-4855-887a-b80d7246e9e9"

	deleteURL := fmt.Sprintf("http://localhost:8080/user/%s", userId)
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
	} else {
		fmt.Println("DELETE request successful")
	}
	fmt.Println(deleteResp)

}
