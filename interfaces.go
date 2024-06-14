package apieasy

import "net/http"

type HTTPHandlerFunc func(ResponseWriter, *http.Request)

type Routes interface {
	ServeHTTP(ResponseWriter, *http.Request)
	Get(string, HTTPHandlerFunc) Router
	Post(string, HTTPHandlerFunc) Router
	Put(string, HTTPHandlerFunc) Router
	Delete(string, HTTPHandlerFunc) Router
	Options(string, HTTPHandlerFunc) Router
	Head(string, HTTPHandlerFunc) Router
	Patch(string, HTTPHandlerFunc) Router
	Run() error
	Handle(method, pattern string, handler HandlerFunc)
}

type ResponseWriter interface {
	WriteHeader(statusCode int)
	Write([]byte) (int, error)
}

func NewRouter(addr ...string) Router {
	return newRouter(addr...)
}
