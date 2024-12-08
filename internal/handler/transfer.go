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

func (h *handler) initTransferGroup(prefix string, router *http.ServeMux) {
	router.Handle(prefix+"/find_all", middleware.Middleware(http.HandlerFunc(h.FindAllTransfer)))
	router.Handle(prefix+"/find_by_id", middleware.Middleware(http.HandlerFunc(h.FindByIdTransfer)))
	router.Handle(prefix+"/find_by_user_id", middleware.Middleware(http.HandlerFunc(h.FindByUserIdTransfer)))
	router.Handle(prefix+"/create", middleware.Middleware(http.HandlerFunc(h.CreateTransfer)))
	router.Handle(prefix+"/update", middleware.Middleware(http.HandlerFunc(h.UpdateTransfer)))
	router.Handle(prefix+"/delete", middleware.Middleware(http.HandlerFunc(h.DeleteTransfer)))
}

func (h *handler) FindAllTransfer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	res, err := h.services.Transfer.FindAll()

	if err != nil {
		response.ResponseError(w, *err)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) FindByIdTransfer(w http.ResponseWriter, r *http.Request) {
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

	res, errRes := h.services.Transfer.FindById(id)

	if err != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) FindByUserIdTransfer(w http.ResponseWriter, r *http.Request) {
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

	res, errRes := h.services.Transfer.FindByUserID(userInt)

	if err != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
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

	var createTransfer requests.CreateTransferRequest

	createTransfer.TransferFrom = userInt

	if err := json.NewDecoder(r.Body).Decode(&createTransfer); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid request",
		}
		response.ResponseError(w, res)
		return
	}

	if err := createTransfer.Validate(); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid validate request",
		}
		response.ResponseError(w, res)
		return
	}

	res, errRes := h.services.Transfer.Create(createTransfer)

	if err != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) UpdateTransfer(w http.ResponseWriter, r *http.Request) {
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

	var updateTransfer requests.UpdateTransferRequest

	updateTransfer.TransferID = id
	updateTransfer.TransferFrom = userInt

	if err := json.NewDecoder(r.Body).Decode(&updateTransfer); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid request",
		}
		response.ResponseError(w, res)
		return
	}

	if err := updateTransfer.Validate(); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid validate request",
		}
		response.ResponseError(w, res)
		return
	}

	res, errErr := h.services.Transfer.Update(updateTransfer)

	if errErr != nil {
		response.ResponseError(w, *errErr)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) DeleteTransfer(w http.ResponseWriter, r *http.Request) {
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

	res, errRes := h.services.Transfer.Delete(id)

	if err != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}
