package apieasy

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Router struct {
	routes     map[string]HTTPHandlerFunc
	handlers   map[string]HandlerFunc
	middleware []func(http.Handler) http.Handler
	addr       string
}

func newRouter(addr ...string) Router {
	var address string
	if len(addr) > 0 {
		address = addr[0]
	} else {
		address = ":8080"
	}

	return Router{
		routes:     make(map[string]HTTPHandlerFunc),
		handlers:   make(map[string]HandlerFunc),
		middleware: []func(http.Handler) http.Handler{},
		addr:       address,
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

func (r *Router) AddMiddleware(mw func(http.Handler) http.Handler) {
	r.middleware = append(r.middleware, mw)
}

func (r *Router) Run() error {
	handler := http.Handler(r)
	for _, mw := range r.middleware {
		handler = mw(handler)
	}

	server := &http.Server{
		Addr:    r.addr,
		Handler: handler,
	}
	return server.ListenAndServe()
}

func (r *Router) Use(mw func(http.Handler) http.Handler) {
	r.middleware = append(r.middleware, mw)
}
