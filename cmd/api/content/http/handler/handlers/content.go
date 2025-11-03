package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/theamandadavis/practice-cms/pkg/id"

	"github.com/go-chi/chi/v5"
)

// GetContentByID retrieves a Content by ID
func (h *ContentHandler) GetContentSummary(w http.ResponseWriter, r *http.Request) {
	contentID := chi.URLParam(r, "id")
	Content, err := h.Service.GetContentSummary(r.Context(), *id.UUIDFromString(&contentID))
	if err != nil {
		http.Error(w, "Content not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(Content)
}
