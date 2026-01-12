package handlers

import (
	"encoding/json"
	"mimiSpikeBall/internal/models"
	"mimiSpikeBall/internal/services"
	"net/http"
)

type CompteHandler struct {
	service *services.CompteService
}

func NewCompteHandler(service *services.CompteService) *CompteHandler {
	return &CompteHandler{service: service}
}

func (h *CompteHandler) ObtenirComptes(w http.ResponseWriter, r *http.Request) {
	comptes, err := h.service.ObtenirComptes()

	if err != nil {
		writeError(w, err, http.StatusBadRequest)
	} else {
		writeSuccess(w, comptes)
	}
}

func (h *CompteHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUserRequest models.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&loginUserRequest)

	if err != nil {
		writeError(w, err, http.StatusBadRequest)
	}

	valid_user, loginUserError := h.service.LoginUser(&loginUserRequest)
	if loginUserError != nil {
		writeError(w, loginUserError, http.StatusBadRequest)
	} else {
		writeSuccess(w, valid_user)
	}
}
