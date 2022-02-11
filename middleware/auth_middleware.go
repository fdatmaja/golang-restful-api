package middleware

import (
	"belajar_golang_restful_api/helper"
	"belajar_golang_restful_api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middle *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		middle.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
