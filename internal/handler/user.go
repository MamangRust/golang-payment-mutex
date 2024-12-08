package handler

import (
	"encoding/json"
	"net/http"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	"payment-mutex/internal/middleware"
	"strconv"
)

func (h *handler) initUserGroup(prefix string, router *http.ServeMux) {
	router.Handle(prefix+"/find_all", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindAllUser)))
	router.Handle(prefix+"/find_by_id", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindByIdUser)))
	router.Handle(prefix+"/create", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.CreateUser)))
	router.Handle(prefix+"/update", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.UpdateUser)))
	router.Handle(prefix+"/delete", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.DeleteUser)))
}

func (h *handler) FindAllUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	res, err := h.services.User.FindAll()

	if err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error find all user",
		}

		response.ResponseError(w, res)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) FindByIdUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error convert id",
		}
		response.ResponseError(w, res)
		return
	}

	res, errRes := h.services.User.FindByID(id)

	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	var createUser requests.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&createUser); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid request",
		}
		response.ResponseError(w, res)
		return
	}

	if err := createUser.Validate(); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid validate request",
		}
		response.ResponseError(w, res)
		return
	}

	res, errRes := h.services.User.Create(createUser)

	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error convert id",
		}
		response.ResponseError(w, res)
		return
	}

	var updateUser requests.UpdateUserRequest

	updateUser.UserID = id

	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid request",
		}
		response.ResponseError(w, res)
		return
	}

	if err := updateUser.Validate(); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid validate request",
		}
		response.ResponseError(w, res)
		return
	}

	res, errRes := h.services.User.Update(updateUser)

	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Invalid ID format",
		}
		response.ResponseError(w, res)
		return
	}

	apiRes, errRes := h.services.Saldo.Delete(id)
	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *apiRes)
}
