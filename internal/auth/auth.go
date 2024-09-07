package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("Unauthenticated")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Unauthenticated")
	}

	if vals[0] != "Bearer" {
		return "", errors.New("Unauthenticated")
	}

	return vals[1], nil
}
