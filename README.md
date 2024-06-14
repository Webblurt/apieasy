# APIEasy

APIEasy is a lightweight web framework for Go that simplifies the creation and handling of HTTP requests. It provides a convenient router and context for building robust web applications with minimal boilerplate code.

### Features
1. HTTP Request Handling: Easily define routes and handle HTTP methods.
2. JSON Response: Simplified methods for returning JSON responses with customizable status codes.
3. Flexible Routing: Supports GET, POST, PUT, DELETE, and other HTTP methods.
4. Middleware Support: Extendable middleware for additional request processing.
5. Error Handling: Built-in mechanisms for handling errors and status codes.

### Installation

To use APIEasy in your Go project, follow these steps:

1. Install Go (if not already installed): [https://golang.org/doc/install](https://golang.org/doc/install)
2. Create a new Go module (if not already created):

   ```bash
   go mod init myapp
   ```

3. Install APIEasy:

    ```bash
    go get github.com/Webblurt/apieasy
    ```

Example Usage
Here's a basic example of setting up and running an HTTP server using APIEasy:

    ```golang
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
    ```

### Routing

Routes and handlers are added using the Handle method of the router. Handlers receive a Context object that contains information about the current request and methods to set the response status and message.

### Additional Information
1. APIEasy supports custom HTTP statuses and provides a convenient interface for request handling.
2. Use SetStatus to set the response status and message within a handler.
3. Visit the Go documentation for more information on the standard library.

### Contributing and Issue Reporting
If you have suggestions for improvements or find a bug, please open an issue on GitHub or submit a pull request. Contributions are welcome!

### License

APIEasy is licensed under the MIT License. See the LICENSE file for more information.
