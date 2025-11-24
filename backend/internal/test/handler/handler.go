package handler

import (
	"encoding/json"
	"linux-docker-web-gui/internal/test/service"
	"net/http"
)

type TestHandler struct {
	service *service.TestService
}

func NewHandler(service *service.TestService) *TestHandler {
	return &TestHandler{service: service}
}

func (h *TestHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /test", h.GetTest)
}

func (h *TestHandler) GetTest(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.GetTest()
	if err != nil {
		http.Error(w, "Failed to get test", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
