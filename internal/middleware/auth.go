package middleware

import (
	"net/http"
	"payment-mutex/internal/domain/response"
	"payment-mutex/pkg/auth"
)

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.Authorization(r)

		if err != nil {
			response.ResponseError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		r = auth.SetContextUserId(r, token)

		next.ServeHTTP(w, r)
	})
}

func MiddlewareAuthAndCors(next http.Handler) http.Handler {
	return MiddlewareLogging(MiddlewareCors(MiddlewareAuth(next)))
}
