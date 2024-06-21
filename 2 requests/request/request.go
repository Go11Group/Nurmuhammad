package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Users struct {
	UserID    string    `json:"userId"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Birthday  string    `json:"birthday"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt,omitempty"`
}
type Client struct {
	http.Client
}

func main() {
	router := gin.Default()
	Client := Client{}

	router.GET("/user:id", Client.GetUser)
	router.POST("/user", Client.PostUser)
	router.PUT("/user:id", Client.PutUser)
	router.DELETE("/user:id", Client.DeleteUser)

	router.Run(":8090")

}

func (cl *Client) GetUser(c *gin.Context) {
	userId := c.Param("id")

	getURL := fmt.Sprintf("http://localhost:8080/user/%s", userId)
	resp, err := http.Get(getURL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	c.JSON(http.StatusOK, string(body))
	fmt.Println(string(body))

}

func (cl *Client) PostUser(c *gin.Context) {
	var user Users
	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}
	users, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/user", bytes.NewBuffer(users))

	if err != nil {
		panic(err)
	}

	resp, err := cl.Do(req)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&user)
	c.JSON(http.StatusCreated, user)
}

func (cl *Client) PutUser(c *gin.Context) {
	userId := c.Param("id")

	putURL := fmt.Sprintf("http://localhost:8080/user/%s", userId)
	var user Users
	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}
	users, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("PUT", putURL, bytes.NewBuffer(users))

	if err != nil {
		panic(err)
	}

	resp, err := cl.Do(req)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&user)
	c.JSON(http.StatusCreated, user)
}

func (cl *Client) DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	deleteURL := fmt.Sprintf("http://localhost:8080/user/%s", userId)
	deleteReq, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	deleteResp, err := cl.Do(deleteReq)
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
