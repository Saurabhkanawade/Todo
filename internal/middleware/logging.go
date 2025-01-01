package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		logrus.Infof("Request:%s %s", request.Method, request.URL.Path)
		next.ServeHTTP(w, request)
	})
}
