package main

import (
	"net/http"

	"github.com/webblurt/apieasy"
)

func main() {
	router := apieasy.NewRouter(":8080")

	router.Handle("GET", "/", func(ctx *apieasy.Context) {
		ctx.Writer.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", router)
}
