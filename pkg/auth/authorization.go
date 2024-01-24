package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

type contextKey string

const UserIDKey contextKey = "userId"

func Authorization(r *http.Request) (string, error) {
	keys := r.URL.Query()
	token := keys.Get("token")

	if token != "" {
		return token, errors.New("token not found")
	}

	bearerToken := r.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		tokenString := strings.Split(bearerToken, " ")[1]

		manager, err := NewManager(viper.GetString("JWT_SECRET"))
		if err != nil {
			fmt.Println("Error loading .env file:", err)
			return "", err
		}

		userId, err := manager.ValidateToken(tokenString)

		if err != nil {
			fmt.Println("Error validating token:", err)
			return "", err
		}

		return userId, nil
	}

	return "", errors.New("invalid token")
}
