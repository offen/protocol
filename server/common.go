package server

import (
	"fmt"
	"net/http"
)

func getUserIdentifier(r *http.Request) (string, error) {
	cookie, err := r.Cookie("user")
	if err != nil && err != http.ErrNoCookie {
		return "", fmt.Errorf("offen: error reading user id from cookie")
	}
	if err == http.ErrNoCookie {
		return "", nil
	}
	return cookie.Value, nil
}
