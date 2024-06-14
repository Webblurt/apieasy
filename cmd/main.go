package main

import (
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
}
