package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	s := headers.Get("Authorization")
	if s == "" {
		return "", errors.New("No auth info found")
	}
	val := strings.Split(s, " ")
	if len(val) != 2 || val[0] != "Apikey" {
		return "", errors.New("Malformend auth header")
	}
	return val[1], nil

}
