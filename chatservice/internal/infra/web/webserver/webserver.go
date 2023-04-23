package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Webserver struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(webServerPort string) *Webserver {
	return &Webserver{
		WebServerPort: webServerPort,
	}
}

func (s *Webserver) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *Webserver) Start() {

	s.Router.Use(middleware.Logger)
	s.Router = chi.NewRouter()

	for path, handler := range s.Handlers {
		s.Router.Get(path, handler)
	}

	if err := http.ListenAndServe(s.WebServerPort, s.Router); err != nil {
		panic(err.Error())
	}
}
