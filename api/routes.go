package api

import (
	"kafka/api/handlers"
	"kafka/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	router   *mux.Router
	handlers *handlers.Handlers
}

func NewRouter(service *service.Service) *Router {
	router := mux.NewRouter()

	handler := handlers.NewHandlers(service)

	router.HandleFunc("/messages", handler.Get).Methods(http.MethodGet)
	router.HandleFunc("/messages", handler.Create).Methods(http.MethodPost)

	return &Router{
		router:   router,
		handlers: handler,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
