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
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	var login requests.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid request",
		}
		response.ResponseError(w, res)
		return
	}

	if err := login.Validate(); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid validate request",
		}
		response.ResponseError(w, res)
		return
	}

	res, err := h.services.Auth.Login(&login)
	if err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error login user",
		}
		response.ResponseError(w, res)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	var register requests.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid request",
		}
		response.ResponseError(w, res)
		return
	}

	if err := register.Validate(); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid validate request",
		}
		response.ResponseError(w, res)
		return
	}

	res, err := h.services.Auth.RegisterUser(&register)
	if err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error creating user",
		}
		response.ResponseError(w, res)
		return
	}

	response.ResponseMessage(w, *res)
}
