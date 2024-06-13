# APIEasy - Simple Go Web Framework

APIEasy is a lightweight web framework for Go that makes it easy to create and handle HTTP requests using a convenient router and context.

### Installation

To use APIEasy in your Go project, follow these steps:

1. Install Go (if not already installed): [https://golang.org/doc/install](https://golang.org/doc/install)
2. Create a new Go module (if not already created):

   ```bash
   go mod init myapp
   ```

3. Install APIEasy:

    ```bash
    go get github.com/webblurt/apieasy
    ```

Example Usage
Here's a basic example of setting up and running an HTTP server using APIEasy:

    ```go
    package main

    import (
        "github.com/webblurt/apieasy"
    )

    func main() {
        router := apieasy.NewRouter(":8080")

        // Adding a handler for a GET request
        router.Handle("GET", "/", func(ctx *apieasy.Context) {
            ctx.SetStatus(apieasy.OK, "Hello, World!")
        })

        // Running the server
        if err := router.Run(); err != nil {
            panic(err)
        }
    }
    ```

Routing
Routes and handlers are added using the Handle method of the router. Handlers receive a Context object that contains information about the current request and methods to set the response status and message.

Additional Information
APIEasy supports custom HTTP statuses and provides a convenient interface for request handling.
Use SetStatus to set the response status and message within a handler.
Visit the Go documentation for more information on the standard library.
Contributing and Issue Reporting
If you have suggestions for improvements or find a bug, please open an issue on GitHub or submit a pull request. Contributions are welcome!

License
APIEasy is licensed under the MIT License. See the LICENSE file for more information.
