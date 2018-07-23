package main

import (
	"net/http"
	"errors"
)

func session(w http.ResponseWriter, r *http.Request)(sess data.Session, err error){
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}