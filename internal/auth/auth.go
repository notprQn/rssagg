package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authetication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 1 {
		return "", errors.New("malformed auth header")
	}
	return vals[0], nil
}
