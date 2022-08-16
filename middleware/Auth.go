package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	//"go.mongodb.org/mongo-driver/x/mongo/driver/session"
)

func Authenticate(hand http.HandlerFunc) http.HandlerFunc{
	return http.HandleFunc(func (w http.ResponseWriter, r *http.Request)  {
		session := SessionHandler.GetSession(r)
		if session == nil{
			w.Write([]byte("UnAuthorized user"))
		}
		hand.ServeHTTP(w,r)

	})
}

//var  y= mux.MiddlewareFunc(http.handler)