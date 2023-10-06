package handlers

import (
	"encoding/json"
	"kafka/models"
	"kafka/service"
	"net/http"
)

type Handlers struct {
	service service.Service
}

func NewHandlers(service *service.Service) *Handlers {
	return &Handlers{
		service: *service,
	}
}

func (h *Handlers) Get(w http.ResponseWriter, req *http.Request) {

}

func (h *Handlers) Create(w http.ResponseWriter, req *http.Request) {
	var newData models.Data
	err := json.NewDecoder(req.Body).Decode(&newData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := h.service.Create(newData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResoponse(w, data)
}

func writeResoponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
