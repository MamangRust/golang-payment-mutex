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

func ResponseMessage(w http.ResponseWriter, message string, data interface{}, status int) {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}

func ResponseError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json") // pastikan header ditetapkan sebelum WriteHeader

	w.WriteHeader(status)

	res := Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal("Error encoding response: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
