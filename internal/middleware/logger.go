package middleware

import (
	"fmt"
	"net/http"
	"time"
)

const (
	greenColor = "\033[32m"
	resetColor = "\033[0m"
)

func MiddlewareLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		defer func() {
			duration := time.Since(startTime)
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			logMessage := fmt.Sprintf("[%s] %s %s %s %s\n", currentTime, r.Method, r.URL.Path, r.RemoteAddr, duration)
			coloredLog := fmt.Sprintf("%s%s%s", greenColor, logMessage, resetColor)
			fmt.Print(coloredLog)
		}()
		next.ServeHTTP(w, r)
	})
}

// func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		startTime := time.Now()

// 		defer func() {
// 			duration := time.Since(startTime)
// 			currentTime := time.Now().Format("2006-01-02 15:04:05")
// 			logMessage := fmt.Sprintf("[%s] %s %s %s %s\n", currentTime, r.Method, r.URL.Path, r.RemoteAddr, duration)
// 			coloredLog := fmt.Sprintf("%s%s%s", greenColor, logMessage, resetColor)
// 			fmt.Print(coloredLog)
// 		}()

// 		next.ServeHTTP(w, r)
// 	}
// }
