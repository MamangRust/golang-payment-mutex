package handler

import (
	"encoding/json"
	"net/http"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	"payment-mutex/internal/middleware"
	"payment-mutex/pkg/auth"
	"strconv"
)

func (h *handler) initTopupGroup(prefix string, router *http.ServeMux) {
	router.Handle(prefix+"/find_all", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindAllTopup)))
	router.Handle(prefix+"/find_by_id", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindByIdTopup)))
	router.Handle(prefix+"/find_by_user_id", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindByUserIdTopup)))
	router.Handle(prefix+"/find_by_users_id", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindByUsersIdTopup)))
	router.Handle(prefix+"/create", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.CreateTopup)))
	router.Handle(prefix+"/update", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.UpdateTopup)))
	router.Handle(prefix+"/delete", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.DeleteTopup)))
}

func (h *handler) FindAllTopup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	res, err := h.services.Topup.FindAll()

	if err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error find all topup",
		}

		response.ResponseError(w, res)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) FindByIdTopup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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

	res, errRes := h.services.Topup.FindById(id)

	if err != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) FindByUsersIdTopup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	userId := auth.GetContextUserId(r)

	userInt, err := strconv.Atoi(userId)

	if err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error convert id",
		}
		response.ResponseError(w, res)
		return
	}

	res, errRes := h.services.Topup.FindByUsersID(userInt)

	if err != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) FindByUserIdTopup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	userId := auth.GetContextUserId(r)

	userInt, err := strconv.Atoi(userId)

	if err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error convert id",
		}
		response.ResponseError(w, res)
		return
	}

	res, errRes := h.services.Topup.FindByUserID(userInt)

	if err != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) CreateTopup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userId := auth.GetContextUserId(r)

	userInt, err := strconv.Atoi(userId)

	if err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error convert id",
		}
		response.ResponseError(w, res)
		return
	}

	var createTopup requests.CreateTopupRequest

	createTopup.UserID = userInt

	if err := json.NewDecoder(r.Body).Decode(&createTopup); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid request",
		}

		response.ResponseError(w, res)

	}

	if err := createTopup.Validate(); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid validate request",
		}

		response.ResponseError(w, res)
		return
	}

	res, errRes := h.services.Topup.Create(createTopup)

	if err != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) UpdateTopup(w http.ResponseWriter, r *http.Request) {
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
	}

	userId := auth.GetContextUserId(r)

	userInt, err := strconv.Atoi(userId)

	if err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error convert user id",
		}
		response.ResponseError(w, res)
		return
	}

	var updateTopup requests.UpdateTopupRequest

	updateTopup.TopupID = id
	updateTopup.UserID = userInt

	if err := json.NewDecoder(r.Body).Decode(&updateTopup); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid request",
		}
		response.ResponseError(w, res)
		return
	}

	if err := updateTopup.Validate(); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid validate request",
		}
		response.ResponseError(w, res)
		return
	}

	res, errRes := h.services.Topup.Update(updateTopup)

	if err != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) DeleteTopup(w http.ResponseWriter, r *http.Request) {
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

	res, errRes := h.services.Topup.Delete(id)

	if err != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}
