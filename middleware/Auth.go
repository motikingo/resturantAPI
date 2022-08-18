package middleware

import (
	"net/http"
	handler "github.com/motikingo/resturant-api/httpHandler"
)


type MiddlewareHandler struct{
	session *handler.SessionHandler
}

func NewMiddlewareHandler(	session *handler.SessionHandler) *MiddlewareHandler  {
	return &MiddlewareHandler{session:session}
}

func (middleHa *MiddlewareHandler) Authenticate(hand http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		sess := middleHa.session.GetSession(r)
		if sess == nil{
			http.Error(w,http.StatusText(http.StatusUnauthorized),http.StatusUnauthorized)
			return
		}
		hand.ServeHTTP(w,r)

	})
}

//var  y= mux.MiddlewareFunc(http.handler)