package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/who", who)          // for query parameters
	http.HandleFunc("/who/", whoWithPath) // for path parameters
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func who(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)

	switch r.Method {
	case "GET":
		handleGet(w, r)
	case "DELETE":
		handleDelete(w, r)
	case "PUT":
		handleUpdate(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("405 - Method Not Allowed"))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	// Read query parameter
	name := r.URL.Query().Get("name")
	response := "I am Nurmuhammad"
	if name != "" {
		response = "Hello, " + name
	}

	_, err := w.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Delete request received"))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("PUT request body:", string(body))

	_, err = w.Write([]byte("Update request received with body: " + string(body)))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}

// Handle path parameters
func whoWithPath(w http.ResponseWriter, r *http.Request) {
	// Extract path parameter
	path := strings.TrimPrefix(r.URL.Path, "/who/")
	if path == "" {
		http.Error(w, "User ID is missing", http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("User ID from path parameter: %s", path)
	_, err := w.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}
