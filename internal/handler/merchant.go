package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"payment-mutex/internal/domain/requests"
	"payment-mutex/internal/domain/response"
	"payment-mutex/internal/middleware"
	"strconv"
)

func (h *handler) InitMerchantGroup(prefix string, router *http.ServeMux) {
	router.Handle(prefix+"/find_all", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindAllMerchant)))
	router.Handle(prefix+"/find_by_id", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindMerchantByID)))
	router.Handle(prefix+"/find_by_name", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.FindMerchantByName)))
	router.Handle(prefix+"/create", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.CreateMerchant)))
	router.Handle(prefix+"/update", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.UpdateMerchant)))
	router.Handle(prefix+"/delete", middleware.MiddlewareAuthAndCors(http.HandlerFunc(h.DeleteMerchant)))

}

func (h *handler) FindAllMerchant(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := r.URL.Query().Get("search")

	res, errRes := h.services.Merchant.FindAll(page, pageSize, search)

	if err != nil {

		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)

}

func (h *handler) FindMerchantByID(w http.ResponseWriter, r *http.Request) {
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

	res, errRes := h.services.Merchant.FindByID(id)

	if errRes != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Error find merchant by id",
		}

		response.ResponseError(w, res)
		return
	}

	response.ResponseMessage(w, *res)

}

func (h *handler) FindMerchantByName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Name parameter is required",
		}
		response.ResponseError(w, res)
		return
	}

	res, err := h.services.Merchant.FindByName(name)
	if err != nil {
		response.ResponseError(w, *err)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) CreateMerchant(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	var req requests.CreateMerchantRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Invalid input",
		}
		fmt.Println(err.Error())
		response.ResponseError(w, res)
		return
	}

	res, err := h.services.Merchant.Create(req)
	if err != nil {
		response.ResponseError(w, *err)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) UpdateMerchant(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	var req requests.UpdateMerchantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Invalid input",
		}

		fmt.Println(err)
		response.ResponseError(w, res)
		return
	}

	res, err := h.services.Merchant.Update(req)
	if err != nil {
		response.ResponseError(w, *err)
		return
	}

	response.ResponseMessage(w, *res)
}

func (h *handler) DeleteMerchant(w http.ResponseWriter, r *http.Request) {
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

	res, errRes := h.services.Merchant.Delete(id)

	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}
