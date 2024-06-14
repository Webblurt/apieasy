package apieasy

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Code    StatusCode
	Message string
}

type Context struct {
	Writer         http.ResponseWriter
	Request        *http.Request
	ResponseStatus Status
	Status         StatusCode
	Message        string
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
	}
}

func (c *Context) JSON(status StatusCode, data interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")

	if status == 0 {
		status = Success.Default
	}

	c.Writer.WriteHeader(int(status))

	if err := json.NewEncoder(c.Writer).Encode(data); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) SetStatus(code StatusCode, message string, err error) {
	c.Status = code
	c.Message = message
	c.ResponseStatus = Status{Code: code, Message: message}
}
