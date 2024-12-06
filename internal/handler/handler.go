package handler

import (
	"net/http"
	"payment-mutex/internal/service"
)

type handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *handler {
	return &handler{
		services: services,
	}
}

func (h *handler) Init() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		w.Write([]byte("Hello World"))
	})

	h.InitApi(router)

	return router
}

func (h *handler) InitApi(r *http.ServeMux) {
	h.initAuthGroup("/auth", r)
	h.initSaldoGroup("/saldo", r)
	h.initTopupGroup("/topup", r)
	h.initTransferGroup("/transfer", r)
	h.InitWithdrawGroup("/withdraw", r)
	h.initUserGroup("/user", r)
}
