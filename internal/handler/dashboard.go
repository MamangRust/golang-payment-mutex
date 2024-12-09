package handler

import (
	"net/http"
	"payment-mutex/internal/domain/response"
	"payment-mutex/internal/middleware"
)

func (h *handler) initDashboardGroup(prefix string, router *http.ServeMux) {
	router.Handle(prefix+"", middleware.Middleware(http.HandlerFunc(h.Dashboard)))
}

func (h *handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := response.ErrorResponse{
			Status:  "error",
			Message: "Method Not Allowed",
		}
		response.ResponseError(w, res)
		return
	}

	res, errRes := h.services.Dashboard.GetGlobalOverview()

	if errRes != nil {
		response.ResponseError(w, *errRes)
		return
	}

	response.ResponseMessage(w, *res)
}
