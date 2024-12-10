package middleware

import (
	"context"
	"net/http"
	"payment-mutex/internal/service"
)

type ApiKeyKey struct{}

func MerchantMiddleware(next http.HandlerFunc, merchantService service.MerchantService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-Api-Key")
		if apiKey == "" {
			http.Error(w, "API Key is required", http.StatusUnauthorized)
			return
		}

		_, err := merchantService.FindByApiKey(apiKey)
		if err != nil {
			http.Error(w, "Invalid API Key", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ApiKeyKey{}, apiKey)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
