package handler

import (
	"net/http"
	"time"
	//"fmt"
	"github.com/golang-jwt/jwt/v4"

	"github.com/motikingo/resturant-api/entity"
)

var key = []byte("this is my key")

//var err error
type SessionHandler struct {
}

func NewSessionHandler() *SessionHandler {
	return &SessionHandler{}
}

func (sessionHand *SessionHandler) CreateSession(session *entity.Session, w http.ResponseWriter) bool {

	expiretime := time.Now().Add(24 * time.Hour)
	session.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: expiretime},
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, session)

	tknstr, er := tkn.SignedString(key)

	if er != nil {
		w.Write([]byte(er.Error()))
		return false
	}

	cookie := http.Cookie{
		Name:       entity.SessionName,
		Value:      tknstr,
		Path:       "/",
		Domain:     "",
		Expires:    expiretime,
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   true,
		SameSite:   0,
		Raw:        "",
		Unparsed:   []string{},
	}

	http.SetCookie(w, &cookie)

	return true

}

func (sessionHand *SessionHandler) GetSession(r *http.Request) *entity.Session {

	cookie, err := r.Cookie(entity.SessionName)
	session := &entity.Session{}
	if err != nil {
		return nil
	}

	tknstr := cookie.Value
	//fmt.Println(tknstr)
	tkn, err := jwt.ParseWithClaims(tknstr, session, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil
	}

	if !tkn.Valid {
		return nil

	}

	return session

}

func (sessionHand *SessionHandler) DeleteSession(w http.ResponseWriter) error {
	var session entity.Session

	expireTime := time.Now().Add(-24 * time.Hour)
	session.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: expireTime},
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, session)

	tknstr, er := tkn.SignedString(key)

	if er != nil {
		return er
	}

	cookie := http.Cookie{
		Name:     entity.SessionName,
		Value:    tknstr,
		HttpOnly: true,
		Expires:  expireTime,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)
	return nil
}
