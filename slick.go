package slick

import (
	"context"
	"github.com/a-h/templ"
	"github.com/julienschmidt/httprouter"
	"log/slog"
	"net/http"
)

type ErrorHandler func(error, *Context) error

type Context struct {
	response http.ResponseWriter
	request  *http.Request
	ctx      context.Context
}

func (c *Context) Render(component templ.Component) error {
	return component.Render(c.ctx, c.response)
}

type Plug func(Handler) Handler

type Handler func(c *Context) error

type Slick struct {
	ErrorHandler ErrorHandler
	router       *httprouter.Router
	middlewares  []Plug
}

func New() *Slick {
	return &Slick{
		ErrorHandler: defaultErrorHandler,
		router:       httprouter.New(),
	}
}

func (s *Slick) Plug(plugs ...Plug) {
	s.middlewares = append(s.middlewares, plugs...)
}

func (s *Slick) Start(port string) error {
	slog.Info("starting server on", "port", port)

	s.router.ServeFiles("/public/*filepath", http.Dir("./src/public"))

	return http.ListenAndServe(port, s.router)
}

func (s *Slick) Get(path string, h Handler, plugs ...Handler) {
	s.router.GET(path, s.makeHTTPRouterHandler(h))
}

func (s *Slick) makeHTTPRouterHandler(h Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ctx := &Context{
			response: w,
			request:  r,
			ctx:      context.Background(),
		}

		for _, mw := range s.middlewares {
			h = mw(h)
		}

		if err := h(ctx); err != nil {
			// TODO: handle the error from the error handlers huh?)
			s.ErrorHandler(err, ctx)
		}
	}
}

func defaultErrorHandler(err error, c *Context) error {
	slog.Error("error", "err", err)
	return nil
}
