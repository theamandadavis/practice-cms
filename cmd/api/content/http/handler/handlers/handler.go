package handlers

import (
	"github.com/go-chi/chi/v5"
	content "github.com/theamandadavis/practice-cms/cmd/content/internal"
)

type ContentHandler struct {
	Service content.Service
}

func NewContentHandler(service *content.Service) *ContentHandler {
	return &ContentHandler{Service: *service}
}

func (h *ContentHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/content-summary/{id}", h.GetContentSummary)

	return r
}
