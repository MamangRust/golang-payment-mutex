package auth

import (
	"context"
	"net/http"
)

func SetContextUserId(r *http.Request, token string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), UserIDKey, token))
}

func GetContextUserId(r *http.Request) string {
	userId := r.Context().Value(UserIDKey)
	if userId == nil {
		return ""
	}
	return userId.(string)
}
