package handler

import (
	"encoding/json"
	"net/http"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	"payment-mutex/internal/middleware"
)

func (h *handler) initAuthGroup(prefix string, router *http.ServeMux) {
	router.Handle(prefix+"/login", middleware.Middleware(http.HandlerFunc(h.login)))
	router.Handle(prefix+"/register", middleware.Middleware(http.HandlerFunc(h.register)))
}

func (h *handler) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var login requests.AuthRequest

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := login.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Auth.Login(&login)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error login user")
		return
	}

	response.ResponseToken(w, "Success login", res, http.StatusOK)

}

func (h *handler) register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
		return
	}

	var register requests.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
		return
	}

	if err := register.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Auth.RegisterUser(&register)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error create user")
		return
	}

	response.ResponseMessage(w, "Success create user", res, http.StatusCreated)

}
