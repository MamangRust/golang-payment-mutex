package handler

import (
	"encoding/json"
	"net/http"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	"payment-mutex/internal/middleware"
	"strconv"
)

func (h *handler) initTransactionGroup(prefix string, router *http.ServeMux) {
	router.Handle(prefix+"/find_all", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindAllTransaction)))
	router.Handle(prefix+"/find_by_id", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindByIdTransaction)))
	router.Handle(prefix+"/create", middleware.MiddlewareAuthAndCors(middleware.MerchantMiddleware(http.HandlerFunc(h.CreateTransaction), h.services.Merchant)))
	router.Handle(prefix+"/update", middleware.MiddlewareAuthAndCors(middleware.MerchantMiddleware(http.HandlerFunc(h.UpdateTransaction), h.services.Merchant)))
	router.Handle(prefix+"/delete", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.DeleteTransaction)))
}

func (h *handler) FindAllTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	res, err := h.services.Transaction.FindAll()

	if err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error find all transaction",
		}

		response.ResponseError(w, res)

		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) FindByIdTransaction(w http.ResponseWriter, r *http.Request) {
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

	res, errRes := h.services.Transaction.FindById(id)
	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	apiKey, ok := r.Context().Value(middleware.ApiKeyKey{}).(string)

	if !ok || apiKey == "" {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized",
		}
		response.ResponseError(w, res)
		return
	}

	var createTransaction requests.CreateTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&createTransaction); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid request",
		}
		response.ResponseError(w, res)
		return
	}

	if err := createTransaction.Validate(); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid validate request",
		}
		response.ResponseError(w, res)
		return
	}

	res, errRes := h.services.Transaction.Create(apiKey, createTransaction)
	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	apiKey, ok := r.Context().Value(middleware.ApiKeyKey{}).(string)
	if !ok || apiKey == "" {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Unauthorized",
		}
		response.ResponseError(w, res)
		return
	}

	var updateTransaction requests.UpdateTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&updateTransaction); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid request",
		}
		response.ResponseError(w, res)
		return
	}

	if err := updateTransaction.Validate(); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid validate request",
		}
		response.ResponseError(w, res)
		return
	}

	res, errRes := h.services.Transaction.Update(apiKey, updateTransaction)
	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
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

	res, errRes := h.services.Transaction.Delete(id)
	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}
