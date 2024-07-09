package apieasy

import (
	"encoding/json"
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type Router struct {
	routes      map[string]HTTPHandlerFunc
	handlers    map[string]HandlerFunc
	middlewares []MiddlewareFunc
	addr        string
}

func newRouter(addr ...string) Router {
	var address string
	if len(addr) > 0 {
		address = addr[0]
	} else {
		address = ":8080"
	}

	return Router{
		routes:      make(map[string]HTTPHandlerFunc),
		handlers:    make(map[string]HandlerFunc),
		middlewares: []MiddlewareFunc{},
		addr:        address,
	}
}

func (r *Router) Handle(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func colorForStatus(status int) string {
	switch {
	case status >= 200 && status < 300:
		return "\033[42m" // Green
	case status >= 300 && status < 400:
		return "\033[44m" // Blue
	case status >= 400 && status < 500:
		return "\033[43m" // Yellow/Orange
	case status >= 500:
		return "\033[41m" // Red
	default:
		return "\033[0m" // Reset
	}
}

func colorForMethod(method string) string {
	switch {
	case method == "GET":
		return "\033[36m" // Light blue
	case method == "POST":
		return "\033[32m" // Green
	case method == "PUT":
		return "\033[33m" // Yellow
	case method == "DELETE":
		return "\033[31m" // Red
	case method == "OPTIONS":
		return "\033[35m" // Purple
	case method == "HEAD":
		return "\033[34m" // Blue
	case method == "PATCH":
		return "\033[37m" // White
	default:
		return "\033[0m" // Reset
	}
}

func resetColor() string {
	return "\033[0m"
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		key := req.Method + "-" + req.URL.Path
		if handler, ok := r.handlers[key]; ok {
			ctx := NewContext(w, req)
			handler(ctx)

			if ctx.Status != 0 {
				w.WriteHeader(int(ctx.Status))
				if ctx.Message != nil {
					switch msg := ctx.Message.(type) {
					case string:
						w.Write([]byte(msg))
					default:
						json.NewEncoder(w).Encode(msg)
					}
				}
				scolor := colorForStatus(int(ctx.Status))
				mcolor := colorForMethod(req.Method)
				reset := resetColor()
				log.Printf("|Controller: |%s%s%s||%s| |%s%d%s| |%v|", mcolor, req.Method, reset, req.URL.Path, scolor, ctx.Status, reset, ctx.Message)
			}
		} else {
			http.NotFound(w, req)
			scolor := colorForStatus(404)
			mcolor := colorForMethod(req.Method)
			reset := resetColor()
			log.Printf(" |Request not found: |%s||%s|| |%s404%s|", mcolor, req.Method, reset, req.URL.Path, scolor, reset)
		}
	})

	for _, mw := range r.middlewares {
		handler = mw(handler)
	}

	handler.ServeHTTP(w, req)
}

func (r *Router) AddMiddleware(mw func(http.Handler) http.Handler) {
	r.middlewares = append(r.middlewares, mw)
}

func (r *Router) Use(mw MiddlewareFunc) {
	r.middlewares = append(r.middlewares, mw)
}

func (r *Router) Run() error {
	server := &http.Server{
		Addr:    r.addr,
		Handler: r,
	}
	return server.ListenAndServe()
}
