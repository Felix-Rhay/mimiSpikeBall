package handlers

import (
	"encoding/json"
	"mimiSpikeBall/internal/models"
	"net/http"
)

func writeJSON[T any](w http.ResponseWriter, status int, payload T) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, err error, status int) {
	writeJSON(w, status, models.AppelSortant[any]{
		Code:    status,
		Message: err.Error(),
	})
}

func writeSuccess[T any](w http.ResponseWriter, data T) {
	writeJSON(w, http.StatusOK, models.AppelSortant[T]{
		Code:   http.StatusOK,
		Sortie: data,
	})
}
