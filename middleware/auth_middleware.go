package middleware

import "net/http"

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}

}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-KEY") {
		//OK
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		//error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		WebResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTORIZED",
		}
		helper.WriteToResponseBody(writer, WebResponse)
	}

}
