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
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	res, err := h.services.Transfer.FindAll()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error find all transfer")
		return
	}

	response.ResponseMessage(w, "Success find all transfer", res, http.StatusOK)
}

func (h *handler) FindByIdTransfer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert id")
	}

	res, err := h.services.Transfer.FindById(id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error find transfer by id")
		return
	}

	response.ResponseMessage(w, "Success find transfer by id", res, http.StatusOK)
}

func (h *handler) FindByUserIdTransfer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userId := auth.GetContextUserId(r)

	userInt, err := strconv.Atoi(userId)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert user id")
	}

	res, err := h.services.Transfer.FindByUserID(userInt)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error find transfer by user id")
		return
	}

	response.ResponseMessage(w, "Success find transfer by user id", res, http.StatusOK)
}

func (h *handler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userId := auth.GetContextUserId(r)

	userInt, err := strconv.Atoi(userId)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert user id")
		return
	}

	var createTransfer requests.CreateTransferRequest

	createTransfer.TransferFrom = userInt

	if err := json.NewDecoder(r.Body).Decode(&createTransfer); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := createTransfer.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Transfer.Create(createTransfer)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error create transfer")
		return
	}

	response.ResponseMessage(w, "Success create transfer", res, http.StatusOK)
}

func (h *handler) UpdateTransfer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert id")
	}

	userId := auth.GetContextUserId(r)

	userInt, err := strconv.Atoi(userId)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert user id")
		return
	}

	var updateTransfer requests.UpdateTransferRequest

	updateTransfer.TransferID = id
	updateTransfer.TransferFrom = userInt

	if err := json.NewDecoder(r.Body).Decode(&updateTransfer); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := updateTransfer.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Transfer.Update(updateTransfer)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error update transfer")
		return
	}

	response.ResponseMessage(w, "Success update transfer", res, http.StatusOK)
}

func (h *handler) DeleteTransfer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert id")
	}

	err = h.services.Transfer.Delete(id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error delete transfer")
		return
	}

	response.ResponseMessage(w, "Success delete transfer", nil, http.StatusOK)
}
