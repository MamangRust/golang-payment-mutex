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
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	res, err := h.services.User.FindAll()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error find all saldo")
		return
	}

	response.ResponseMessage(w, "Success find all user", res, http.StatusOK)
}

func (h *handler) FindByIdUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert id")
		return
	}

	res, err := h.services.User.FindByID(id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error find saldo by user id")
		return
	}

	response.ResponseMessage(w, "Success find user by user id", res, http.StatusOK)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var createUser requests.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&createUser); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := createUser.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.User.Create(createUser)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error create saldo")
		return
	}

	response.ResponseMessage(w, "Success create saldo", res, http.StatusCreated)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert id")
	}

	var updateUser requests.UpdateUserRequest

	updateUser.UserID = id

	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := updateUser.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.User.Update(updateUser)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error update saldo")
		return
	}

	response.ResponseMessage(w, "Success update user", res, http.StatusOK)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
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
