package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Webblurt/apieasy"
)

func main() {
	router := apieasy.NewRouter(":8080")

	router.Handle("GET", "/hello", func(ctx *apieasy.Context) {
		ctx.SetStatus(apieasy.OK, "Hello, World!", nil)
	})

	router.Handle("POST", "/api/data", func(ctx *apieasy.Context) {
		data := struct {
			Message string `json:"message"`
		}{
			Message: "Data received successfully",
		}
		ctx.JSON(apieasy.OK, data)
	})

	if err := router.Run(); err != nil {
		panic(err)
	}

	// Sending request example
	config := apieasy.HTTPClientConfig{
		Timeout:       10 * time.Second,
		SkipTLSVerify: true,
	}

	client := apieasy.NewHTTPClient(config)

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer your-token",
	}

	params := map[string]string{
		"query": "golang",
	}

	url, err := apieasy.AddURLParams("http://example.com/search", params)
	if err != nil {
		fmt.Println("Error adding URL parameters:", err)
		return
	}

	// GET
	response, err := apieasy.SendRequest(client, "GET", url, headers, nil)
	if err != nil {
		fmt.Println("GET Request Error:", err)
	} else {
		fmt.Println("GET Response:", response)
	}

	// POST
	body := strings.NewReader(`{"key": "value"}`)
	response, err = apieasy.SendRequest(client, "POST", "https://example.com/api", headers, body)
	if err != nil {
		fmt.Println("POST Request Error:", err)
	} else {
		fmt.Println("POST Response:", response)
	}

	// PUT
	body = strings.NewReader(`{"key": "updatedValue"}`)
	response, err = apieasy.SendRequest(client, "PUT", "https://example.com/api/1", headers, body)
	if err != nil {
		fmt.Println("PUT Request Error:", err)
	} else {
		fmt.Println("PUT Response:", response)
	}

	// DELETE
	response, err = apieasy.SendRequest(client, "DELETE", "https://example.com/api/1", headers, nil)
	if err != nil {
		fmt.Println("DELETE Request Error:", err)
	} else {
		fmt.Println("DELETE Response:", response)
	}
}
