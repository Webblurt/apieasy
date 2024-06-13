package apieasy

import (
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Status  StatusCode
	Message string
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
	}
}

func (c *Context) JSON(status int, data interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
}

func (c *Context) SetStatus(status StatusCode, message string) {
	c.Status = status
	c.Message = message
}
