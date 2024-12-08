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
	r.Handle(prefix+"/find_by_users_id", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindByUsersIdWithdraw)))
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
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error find all saldo",
		}

		response.ResponseError(w, res)
		return
	}

	response.ResponseMessage(w, *res)

}

func (h *handler) FindByIdWithdraw(w http.ResponseWriter, r *http.Request) {
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

	res, errRes := h.services.Withdraw.FindById(id)

	if err != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) FindByUserIdWithdraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	userId := auth.GetContextUserId(r)

	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error convert user id",
		}
		response.ResponseError(w, res)
		return
	}

	res, errRes := h.services.Withdraw.FindByUserID(userIdInt)

	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) FindByUsersIdWithdraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	userId := auth.GetContextUserId(r)

	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error convert user id",
		}
		response.ResponseError(w, res)
		return
	}

	res, errRes := h.services.Withdraw.FindByUsersID(userIdInt)

	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) CreateWithdraw(w http.ResponseWriter, r *http.Request) {
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

	var request requests.CreateWithdrawRequest

	request.UserID = userInt

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid request",
		}
		response.ResponseError(w, res)
		return
	}

	if err := request.Validate(); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid validate request",
		}
		response.ResponseError(w, res)
		return
	}

	res, errRes := h.services.Withdraw.Create(request)

	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) UpdateWithdraw(w http.ResponseWriter, r *http.Request) {
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

	var updateWithdraw requests.UpdateWithdrawRequest

	updateWithdraw.WithdrawID = id
	updateWithdraw.UserID = userInt

	if err := json.NewDecoder(r.Body).Decode(&updateWithdraw); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid request",
		}
		response.ResponseError(w, res)
		return
	}

	if err := updateWithdraw.Validate(); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error invalid validate request",
		}
		response.ResponseError(w, res)
		return
	}

	res, errRes := h.services.Withdraw.Update(updateWithdraw)

	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) DeleteWithdraw(w http.ResponseWriter, r *http.Request) {
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

	res, errRes := h.services.Withdraw.Delete(id)

	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}
