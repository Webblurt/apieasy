package apieasy

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Router struct {
	handlers map[string]HandlerFunc
	addr     string
}

func NewRouter(addr string) *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
		addr:     addr,
	}
}

func (r *Router) Handle(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := r.handlers[key]; ok {
		ctx := NewContext(w, req)
		handler(ctx)

		if ctx.Status != 0 {
			w.WriteHeader(int(ctx.Status))
			if ctx.Message != "" {
				w.Write([]byte(ctx.Message))
			}
		}
	} else {
		http.NotFound(w, req)
	}
}

func (r *Router) Run() error {
	server := &http.Server{
		Addr:    r.addr,
		Handler: r,
	}
	return server.ListenAndServe()
}
