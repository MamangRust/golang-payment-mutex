package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseAuth struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func ResponseToken(w http.ResponseWriter, message string, token string, status int) {
	res := ResponseAuth{
		Status:  status,
		Message: message,
		Token:   token,
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Fatal(err)
	}
}

func ResponseMessage[T any](w http.ResponseWriter, res ApiResponse[T]) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Fatal("Error encoding response: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func ResponseError(w http.ResponseWriter, res ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Fatal("Error encoding response: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
