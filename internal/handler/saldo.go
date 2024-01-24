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
	router.Handle(prefix+"/create", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.CreateSaldo)))
	router.Handle(prefix+"/update", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.UpdateSaldo)))
	router.Handle(prefix+"/delete", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.DeleteSaldo)))
}

func (h *handler) FindAllSaldo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	res, err := h.services.Saldo.FindAll()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error find all saldo")
		return
	}

	response.ResponseMessage(w, "Success find all saldo", res, http.StatusOK)
}

func (h *handler) FindByIdSaldo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert id")
		return
	}

	res, err := h.services.Saldo.FindById(id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error find saldo by id")
		return
	}

	response.ResponseMessage(w, "Success find saldo by id", res, http.StatusOK)
}

func (h *handler) FindByUserIdSaldo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userId := auth.GetContextUserId(r)

	userInt, err := strconv.Atoi(userId)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert user id")
		return
	}

	res, err := h.services.Saldo.FindByUserID(userInt)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error find saldo by user id")
		return
	}

	response.ResponseMessage(w, "Success find saldo by user id", res, http.StatusOK)
}

func (h *handler) CreateSaldo(w http.ResponseWriter, r *http.Request) {
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

	var createSaldo requests.CreateSaldoRequest

	createSaldo.UserID = userInt

	if err := json.NewDecoder(r.Body).Decode(&createSaldo); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := createSaldo.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Saldo.Create(createSaldo)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error create saldo")
		return
	}

	response.ResponseMessage(w, "Success create saldo", res, http.StatusCreated)
}

func (h *handler) UpdateSaldo(w http.ResponseWriter, r *http.Request) {
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

	var updateSaldo requests.UpdateSaldoRequest

	updateSaldo.UserID = userInt
	updateSaldo.SaldoID = id

	if err := json.NewDecoder(r.Body).Decode(&updateSaldo); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := updateSaldo.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Saldo.Update(updateSaldo)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error update saldo")
		return
	}

	response.ResponseMessage(w, "Success update saldo", res, http.StatusOK)
}

func (h *handler) DeleteSaldo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert id")
	}

	err = h.services.Saldo.Delete(id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error delete saldo")
		return
	}

	response.ResponseMessage(w, "Success delete saldo", nil, http.StatusOK)
}
