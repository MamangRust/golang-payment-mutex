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

func (h *handler) InitWithdrawGroup(prefix string, r *http.ServeMux) {
	r.Handle(prefix+"/find_all", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindAllWithdraw)))
	r.Handle(prefix+"/find_by_id", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindByIdWithdraw)))
	r.Handle(prefix+"/find_by_user_id", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindByUserIdWithdraw)))
	r.Handle(prefix+"/create", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.CreateWithdraw)))
	r.Handle(prefix+"/update", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.UpdateWithdraw)))
	r.Handle(prefix+"/delete", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.DeleteWithdraw)))
}

func (h *handler) FindAllWithdraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	res, err := h.services.Withdraw.FindAll()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error find all withdraw")
		return
	}

	response.ResponseMessage(w, "Success find all withdraw", res, http.StatusOK)

}

func (h *handler) FindByIdWithdraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert id")
	}

	res, err := h.services.Withdraw.FindById(id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error find withdraw by id")
		return
	}

	response.ResponseMessage(w, "Success find withdraw by id", res, http.StatusOK)
}

func (h *handler) FindByUserIdWithdraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userId := auth.GetContextUserId(r)

	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert user id")
	}

	res, err := h.services.Withdraw.FindByUserID(userIdInt)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error find withdraw by user id")
	}

	response.ResponseMessage(w, "Success find withdraw by user id", res, http.StatusOK)
}

func (h *handler) CreateWithdraw(w http.ResponseWriter, r *http.Request) {
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

	var request requests.CreateWithdrawRequest

	request.UserID = userInt

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := request.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Withdraw.Create(request)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error create withdraw")
		return
	}

	response.ResponseMessage(w, "Success create withdraw", res, http.StatusOK)
}

func (h *handler) UpdateWithdraw(w http.ResponseWriter, r *http.Request) {
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

	var updateWithdraw requests.UpdateWithdrawRequest

	updateWithdraw.WithdrawID = id
	updateWithdraw.UserID = userInt

	if err := json.NewDecoder(r.Body).Decode(&updateWithdraw); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := updateWithdraw.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Withdraw.Update(updateWithdraw)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error update withdraw")
		return
	}

	response.ResponseMessage(w, "Success update withdraw", res, http.StatusOK)
}

func (h *handler) DeleteWithdraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error convert id")
	}

	err = h.services.Withdraw.Delete(id)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error delete withdraw")
		return
	}

	response.ResponseMessage(w, "Success delete withdraw", nil, http.StatusOK)
}
