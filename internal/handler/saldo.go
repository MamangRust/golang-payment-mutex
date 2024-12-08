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

func (h *handler) initSaldoGroup(prefix string, router *http.ServeMux) {
	router.Handle(prefix+"/find_all", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindAllSaldo)))
	router.Handle(prefix+"/find_by_id", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindByIdSaldo)))
	router.Handle(prefix+"/find_by_user_id", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindByUserIdSaldo)))
	router.Handle(prefix+"/find_by_users_id", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindByUsersIdSaldo)))
	router.Handle(prefix+"/create", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.CreateSaldo)))
	router.Handle(prefix+"/update", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.UpdateSaldo)))
	router.Handle(prefix+"/delete", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.DeleteSaldo)))
}

func (h *handler) FindAllSaldo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	res, err := h.services.Saldo.FindAll()

	if err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error find all saldo",
		}

		response.ResponseError(w, res)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) FindByIdSaldo(w http.ResponseWriter, r *http.Request) {
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

	apiRes, errRes := h.services.Saldo.FindById(id)
	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *apiRes)
}

func (h *handler) FindByUserIdSaldo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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
			Message: "Error convert user id",
		}
		response.ResponseError(w, res)
		return
	}

	resData, errRes := h.services.Saldo.FindByUserID(userInt)

	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *resData)
}

func (h *handler) FindByUsersIdSaldo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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
			Message: "Error convert user id",
		}
		response.ResponseError(w, res)
		return
	}

	apiRes, errRes := h.services.Saldo.FindByUsersID(userInt)
	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *apiRes)
}

func (h *handler) CreateSaldo(w http.ResponseWriter, r *http.Request) {
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
			Message: "Error convert user id",
		}
		response.ResponseError(w, res)
		return
	}

	var createSaldo requests.CreateSaldoRequest
	createSaldo.UserID = userInt

	if err := json.NewDecoder(r.Body).Decode(&createSaldo); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid request",
		}
		response.ResponseError(w, res)
		return
	}

	if err := createSaldo.Validate(); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid validate request",
		}
		response.ResponseError(w, res)
		return
	}

	apiRes, errRes := h.services.Saldo.Create(createSaldo)
	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *apiRes)
}

func (h *handler) UpdateSaldo(w http.ResponseWriter, r *http.Request) {
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

	var updateSaldo requests.UpdateSaldoRequest
	updateSaldo.UserID = userInt
	updateSaldo.SaldoID = id

	if err := json.NewDecoder(r.Body).Decode(&updateSaldo); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid request",
		}
		response.ResponseError(w, res)
		return
	}

	if err := updateSaldo.Validate(); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid validate request",
		}
		response.ResponseError(w, res)
		return
	}

	apiRes, errRes := h.services.Saldo.Update(updateSaldo)
	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *apiRes)
}

func (h *handler) DeleteSaldo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
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
