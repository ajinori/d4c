package text

import (
	"net/http"
)

type Context interface {
	String(int, string) error
}

func Module(c Context, message string) error {
	return c.String(http.StatusOK, message)
}
