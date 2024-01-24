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
	router.Handle(prefix+"/create", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.CreateTopup)))
	router.Handle(prefix+"/update", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.UpdateTopup)))
	router.Handle(prefix+"/delete", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.DeleteTopup)))
}

func (h *handler) FindAllTopup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	res, err := h.services.Topup.FindAll()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error find all topup")
		return
	}

	response.ResponseMessage(w, "Success find all topup", res, http.StatusOK)
}

func (h *handler) FindByIdTopup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert id")
		return
	}

	res, err := h.services.Topup.FindById(id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error find topup by id")
		return
	}

	response.ResponseMessage(w, "Success find topup by id", res, http.StatusOK)
}

func (h *handler) FindByUserIdTopup(w http.ResponseWriter, r *http.Request) {
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

	res, err := h.services.Topup.FindByUserID(userInt)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error find topup by user id")
		return
	}

	response.ResponseMessage(w, "Success find topup by user id", res, http.StatusOK)
}

func (h *handler) CreateTopup(w http.ResponseWriter, r *http.Request) {
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

	var createTopup requests.CreateTopupRequest

	createTopup.UserID = userInt

	if err := json.NewDecoder(r.Body).Decode(&createTopup); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")

	}

	if err := createTopup.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Topup.Create(createTopup)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error create topup")
		return
	}

	response.ResponseMessage(w, "Success create topup", res, http.StatusCreated)
}

func (h *handler) UpdateTopup(w http.ResponseWriter, r *http.Request) {
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

	var updateTopup requests.UpdateTopupRequest

	updateTopup.TopupID = id
	updateTopup.UserID = userInt

	if err := json.NewDecoder(r.Body).Decode(&updateTopup); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := updateTopup.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Topup.Update(updateTopup)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error update topup")
		return
	}

	response.ResponseMessage(w, "Success update topup", res, http.StatusOK)
}

func (h *handler) DeleteTopup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert id")
	}

	err = h.services.Topup.Delete(id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error delete topup")
		return
	}

	response.ResponseMessage(w, "Success delete topup", nil, http.StatusOK)
}
